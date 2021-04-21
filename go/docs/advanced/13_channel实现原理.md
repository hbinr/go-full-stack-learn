# channel 实现原理

## channel底层数据结构
```go
type hchan struct {
    qcount   uint           // 当前队列中剩余元素个数
    dataqsiz uint           // 环形队列长度，即可以存放的元素个数
    buf      unsafe.Pointer // 环形队列指针
    elemsize uint16         // 每个元素的大小
    closed   uint32            // 标识关闭状态
    elemtype *_type         // 元素类型
    sendx    uint           // 队列下标，指示元素写入时存放到队列中的位置
    recvx    uint           // 队列下标，指示元素从队列的该位置读出
    recvq    waitq          // 等待读消息的goroutine队列
    sendq    waitq          // 等待写消息的goroutine队列
    lock mutex              // 互斥锁，chan不允许并发读写
}
```

从数据结构可以看出channel由队列、类型信息、goroutine等待队列组成，下面分别说明其原理。
### 1. 环形队列
chan内部实现了一个环形队列作为其缓冲区，队列的长度是创建chan时指定的。

下图展示了一个可缓存6个元素的channel示意图：

![](https://static.sitestack.cn/projects/GoExpertProgramming/chapter01/images/chan-01-circle_queue.png)

- `dataqsiz`指示了队列长度为6，即可缓存6个元素；
- `buf`指向队列的内存，队列中还剩余两个元素；
- `qcount`表示队列中还有两个元素；
- `sendx`指示后续写入的数据存储的位置，取值[0, 6)；
- `recvx`指示从该位置读取数据, 取值[0, 6)；

使用数组实现队列是比较常见的操作，sendx表示队首，指示数据写入的位置；recvx表示队尾，指示数据读取的位置
### 2. 等待队列
从channel读数据，如果channel缓冲区为空或者没有缓冲区，当前goroutine会被阻塞，并被加入`recvq`队列 。

向channel写数据，如果channel缓冲区已满或者没有缓冲区，当前goroutine会被阻塞，并被加入`sendq`队列 。

下图展示了一个没有缓冲区的channel，有几个goroutine阻塞等待读数据：

![](https://static.sitestack.cn/projects/GoExpertProgramming/chapter01/images/chan-02-wait_queue.png)

- 因读阻塞的goroutine会被向channel写入数据的goroutine唤醒；
- 因写阻塞的goroutine会被从channel读数据的goroutine唤醒；
**注意:**
>  一般情况下recvq和sendq至少有一个为空。只有一个例外，那就是同一个goroutine使用select语句向channel一边写数据，一边读数据。
### 3. 类型信息
一个channel只能传递一种类型的值，类型信息存储在hchan数据结构中。

- `elemtype`代表类型，用于数据传递过程中的赋值；
- `elemsize`代表类型大小，用于在buf中定位元素位置。
### 4. 互斥锁
一个channel同时仅允许被一个goroutine读写

## channel 操作
### 1. 创建channel
创建channel的过程实际上是初始化`hchan`结构。其中类型信息和缓冲区长度由`make`语句传入，`buf`的大小则与元素大小和缓冲区长度共同决定。

创建channel的伪代码如下所示：
```go
func makechan(t *chantype, size int) *hchan {
    var c *hchan
    c = new(hchan)
    c.buf = malloc(元素类型大小*size)
    c.elemsize = 元素类型大小
    c.elemtype = 元素类型
    c.dataqsiz = size
    return c
}
```

### 2. 向channel写数据
向一个channel中写数据简单过程如下：

1. 如果等待接收队列recvq不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从recvq取出G, 并把数据写入，最后把该G唤醒，结束发送过程；
2. 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程；
3. 如果缓冲区中没有空余位置，将待发送数据写入G，将当前G加入sendq，进入睡眠，等待被读goroutine唤醒；
   
简单流程图如下：
![](https://static.sitestack.cn/projects/GoExpertProgramming/chapter01/images/chan-03-send_data.png)
### 3. 从channel读数据 
从一个channel读数据简单过程如下：

1. 如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束读取过程；
2. 如果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾部，把G唤醒，结束读取过程；
3. 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程；
4. 将当前goroutine加入recvq，进入睡眠，等待被写goroutine唤醒；
5. 
简单流程图如下：
![](https://static.sitestack.cn/projects/GoExpertProgramming/chapter01/images/chan-04-recieve_data.png)
### 4. 关闭channel
关闭channel时会把recvq中的G全部唤醒，本该写入G的数据位置为nil。把sendq中的G全部唤醒，但这些G会panic。

除此之外，panic出现的常见场景还有：

1. 关闭值为nil的channel
2. 关闭已经被关闭的channel
3. 向已经关闭的channel写数据

## 常见用法

### 1. 单向channel
顾名思义，单向channel只用用于发送或接收数据

由channel的数据结构我们知道，实际上并没有单向channel
### 2. select 

使用 `select` 可以监控channel，当其中某一个channel可操作时就会触发相应的 `case` 分支。

需要注意：
- `select`语句的多个 `case` 语句的执行顺序是随机的
- `select`的 `case` 语句读管道时不会阻塞，尽管管道中没有数据。这是由于 `case` 语句编译后调用读管道时会明确传入不阻塞的参数，读不到数据时不会将当前协程加入等待队列(recvq)，而是直接返回

### 3. for-range

可以通过 `for-range`  持续地从管道中读出数据，类似遍历一个数组。
range能够感知channel的关闭，当channel被发送数据的协程关闭时，range就会结束，接着退出for循环。

- 当管道中没有数据时会阻塞当前协程，与读管道时的阻塞处理机制一样
- 即使管道被关闭，`for-range` 也可以优雅的结束

