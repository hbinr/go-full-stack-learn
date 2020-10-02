# Go 连接 redis

## go-redis 库

区别于另一个比较常用的 Go 语言 redis client 库：[redigo](https://github.com/gomodule/redigo)，我们这里采用https://github.com/go-redis/redis 连接 Redis 数据库并进行操作，因为 go-redis 支持连接哨兵及集群模式的 Redis。

使用以下命令下载并安装:

```go
go get -u github.com/go-redis/redis
```

## 连接

### 普通连接

```go

package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var RDB *redis.Client

// InitRedis 初始化连接
func InitRedis() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10, // 连接池
	})
	if _, err = RDB.Ping().Result(); err != nil {
		fmt.Println("RDB.Ping() failed,err:", err)
		return err
	}
	return nil
}
```

### 连接 Redis 哨兵模式

```go
func initClient()(err error){
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
```

### 连接 Redis 集群

```go
func initClient()(err error){
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
```
