# Watch

在某些场景下，我们除了要使用 `MULTI/EXEC` 命令外，还需要配合使用 `WATCH` 命令。在用户使用 `WATCH` 命令监视某个键之后，直到该用户执行 `EXEC` 命令的这段时间里，如果有其他用户抢先对被监视的键进行了替换、更新、删除等操作，那么当用户尝试执行 `EXEC` 的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。

Watch 方法接收一个函数和一个或多个 key 作为参数。

```go
Watch(fn func(*Tx) error, keys ...string) error
```

基本使用示例如下，[详细代码](../../code/basic_use/watch/main.go)：

```go
// WatchExample watch 示例：监视watch_count的值，并在值不变的前提下将其值+1
func WatchExample() {
	key := "watch_count"
	err := rdb.RDB.Watch(func(t *redis.Tx) error {
		n, err := t.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 使用Pipeline来执行命令
		_, err = t.Pipelined(func(p redis.Pipeliner) error {
			p.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("WatchExample finished.....")
}
```

参考：

https://www.liwenzhou.com/posts/Go/go_redis/#autoid-2-3-3

更多查阅[go-redis 官方 API](https://godoc.org/github.com/go-redis/redis)
