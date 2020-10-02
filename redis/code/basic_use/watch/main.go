package main

import (
	"fmt"
	"go-full-stack-learn/redis/code/rdb"

	"github.com/go-redis/redis"
)

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

func main() {
	if err := rdb.InitRedis(); err != nil {
		fmt.Println("rdb.InitRedis failed, err:", err)
		return
	}
	defer rdb.RDB.Close()
	WatchExample()
}
