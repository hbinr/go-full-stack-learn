## 一、背景

Pipeline 主要是一种网络优化。它本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。这些命令不能保证在事务中执行。这样做的好处是节省了每个命令的网络往返时间（RTT）。

### 需求：

redis 通过 tcp 来对外提供服务，client 通过 socket 连接发起请求，每个请求在命令发出后会阻塞等待 redis 服务器进行处理，处理完毕后将结果返回给 client。

其实和一个 http 的服务器类似，一问一答，请求一次给一次响应。而这个过程在排除掉 redis 服务本身做复杂操作时的耗时的话，可以看到最耗时的就是这个网络传输过程。每一个命令都对应了发送、接收两个网络传输，假如一个流程需要 0.1 秒，那么一秒最多只能处理 10 个请求，将严重制约 redis 的性能。

在很多场景下，我们要完成一个业务，可能会对 redis 做连续的多个操作，譬如库存减一、订单加一、余额扣减等等，这有很多个步骤是需要依次连续执行的。

### 潜在隐患：

这样的场景，网络传输的耗时将是限制 redis 处理量的主要瓶颈。循环 key，获取 value，可能会造成连接池的连接数增多，连接的创建和摧毁，消耗性能

### 解决方法：

可以引入 pipeline 了，pipeline 管道就是解决执行大量命令时、会产生大量网络请求次数而导致延迟的技术。

其实原理很简单，pipeline 就是把所有的命令一次发过去，避免频繁的发送、接收带来的网络开销，redis 在打包接收到的一堆命令后，依次执行，然后把结果再打包返回给客户端。

### 为什么 Pipeline 这么快？

**先看看原来的多条命令，是如何执行的：**

- Redis Client->>Redis Server: 发送第 1 个命令
- Redis Server->>Redis Client: 响应第 1 个命令
- Redis Client->>Redis Server: 发送第 2 个命令
- Redis Server->>Redis Client: 响应第 2 个命令
- Redis Client->>Redis Server: 发送第 n 个命令
- Redis Server->>Redis Client: 响应第 n 个命令

**Pipeline 机制是怎样的呢：**

- Redis Client->>Redis Server: 发送第 1 个命令（缓存在 Redis Client，未即时发送）
- Redis Client->>Redis Server: 发送第 2 个命令（缓存在 Redis Client，未即时发送）
- Redis Client->>Redis Server: 发送第 n 个命令（缓存在 Redis Client，未即时发送）
- Redis Client->>Redis Server: 发送累积的命令
- Redis Server->>Redis Client: 响应第 1、2、n 个命令

### 基于其特性，Pipepine 有两个明显的局限性：

- 鉴于 Pipepine 发送命令的特性，Redis 服务器是以队列来存储准备执行的命令，而队列是存放在有限的内存中的，所以不宜一次性发送过多的命令。如果需要大量的命令，可分批进行，效率不会相差太远滴，总好过内存溢出

- 由于 pipeline 的原理是收集需执行的命令，到最后才一次性执行。所以无法在中途立即查得数据的结果（需待 pipelining 完毕后才能查得结果），这样会使得无法立即查得数据进行条件判断（比如判断是非继续插入记录）。

## 二、Go 操作 Pipeline

### redis pipeline 示例 1，将命令分开写，简单易理解

```go
// PipelineExample1 redis pipeline示例1，将命令分开写
func PipelineExample1() {
	var incr *redis.IntCmd

	pipe := rdb.RDB.Pipeline()
	incr = pipe.Incr("pipeline_counter1")       // 将key为pipeline_counter的数自增1
	pipe.Expire("pipeline_counter1", time.Hour) // 设置过期时间

	// 执行命令
	if _, err := pipe.Exec(); err != nil {
		fmt.Println("pipe.Exec() failed,err:", err)
		return
	}
	fmt.Println("pipeline_counter1's value:", incr.Val())
}
```

上面的代码相当于将以下两个命令一次发给 redis server 端执行，与不使用 Pipeline 相比能减少一次 RTT。

```sh
INCR pipeline_counter
EXPIRE pipeline_counts 3600
```

### pipeline 示例 2，将命令封装到一个匿名函数中

```go
// PipelineExample2 redis pipeline示例2，将命令封装到一个匿名函数中
func PipelineExample2() {
	var incr *redis.IntCmd
	// 使用Pipelined将命令都封装到一个函数中,声省掉了 Exec()步骤
	_, err := rdb.RDB.Pipelined(func(p redis.Pipeliner) error {
		incr = p.Incr("pipeline_counter2")       // 将key为pipeline_counter的数自增1
		p.Expire("pipeline_counter2", time.Hour) // 设置过期时间
		return nil
	})
	if err != nil {
		fmt.Println("rdb.RDB.Pipelined() failed,err:", err)
		return
	}
	fmt.Println("pipeline_counter2's value:", incr.Val())
}
```

查看 `Pipelined()`底层源码，发现已经做了命令执行操作 `c.Exec()`：

```go

func (c *Client) Pipelined(fn func(Pipeliner) error) ([]Cmder, error) {
	return c.Pipeline().Pipelined(fn)
}

//...

func (c *Pipeline) pipelined(fn func(Pipeliner) error) ([]Cmder, error) {
	if err := fn(c); err != nil {
		return nil, err
    }
    // 执行命令
	cmds, err := c.Exec()
	_ = c.Close()
	return cmds, err
}
```

**注意:**

- 两个示例中的 `incr` 都是预定义好的,这样才能保证在 `pipe.Exec()`执行后，`incr.Val()` 得到正确的返回值。如果直接 `pipe.Incr("pipeline_counter1").Val()`来获取值的话，结果始终是零值，因为 pipeline 中的命令未执行

- 示例 2 中，匿名函数的入参已经封装了 Pipeline 对象，所以不需要再定义 `pipe`，并且 Pipelined()底层已经做了 `Exec()`操作

参考：

https://www.liwenzhou.com/posts/Go/go_redis/#autoid-2-3-0

https://my.oschina.net/u/3266761/blog/3023454

https://www.cnblogs.com/weixuqin/p/10961524.html
