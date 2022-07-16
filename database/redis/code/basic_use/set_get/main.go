package main

import (
	"fmt"

	"hb.study/database/redis/code/rdb"

	"github.com/go-redis/redis"
)

// SetGetExample redis set get操作示例
func SetGetExample() {
	if err := rdb.RDB.Set("httprouter_test", "hello world", 0).Err(); err != nil {
		fmt.Println("set key 'httprouter_test' failed,err:", err)
		return
	}

	val, err := rdb.RDB.Get("httprouter_test").Result()
	if err != redis.Nil {
		fmt.Println("the key 'httprouter_test' dosen't exist,err:", err)
		return
	} else if err != nil {
		fmt.Println("Get httprouter_test'value failed,err:", err)
		return
	}
	fmt.Println("httprouter_test'value is ", val)
}

func main() {
	if err := rdb.InitRedis(); err != nil {
		fmt.Println("rdb.InitRedis failed, err:", err)
		return
	}
	defer rdb.RDB.Close()

	SetGetExample()
}
