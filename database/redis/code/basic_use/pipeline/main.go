package main

import (
	"fmt"
	"time"

	"hb.study/database/redis/code/rdb"

	"github.com/go-redis/redis"
)

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

func main() {
	if err := rdb.InitRedis(); err != nil {
		fmt.Println("rdb.InitRedis failed, err:", err)
		return
	}
	defer rdb.RDB.Close()
	PipelineExample1()
	PipelineExample2()
}
