package main

import (
	"fmt"
	"go-full-stack-learn/redis/code/rdb"

	"github.com/go-redis/redis"
)

// SetGetExample redis set get操作示例
func SetGetExample() {
	if err := rdb.RDB.Set("test", "hello world", 0).Err(); err != nil {
		fmt.Println("set key 'test' failed,err:", err)
		return
	}

	val, err := rdb.RDB.Get("test").Result()
	if err != redis.Nil {
		fmt.Println("the key 'test' dosen't exist,err:", err)
		return
	} else if err != nil {
		fmt.Println("Get test'value failed,err:", err)
		return
	}
	fmt.Println("test'value is ", val)
}

func main() {
	if err := rdb.InitRedis(); err != nil {
		fmt.Println("rdb.InitRedis failed, err:", err)
		return
	}
	defer rdb.RDB.Close()

	SetGetExample()
}
