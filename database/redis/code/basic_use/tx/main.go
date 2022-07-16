package main

import (
	"fmt"
	"time"

	"hb.study/database/redis/code/rdb"

	"github.com/go-redis/redis"
)

// TxExample1 redis 事务操作示例1
func TxExample1() {
	var incr *redis.IntCmd
	pipe := rdb.RDB.TxPipeline()
	incr = pipe.Incr("tx_pipeline_counter1")       // 将key为pipeline_counter的数自增1
	pipe.Expire("tx_pipeline_counter1", time.Hour) // 设置过期时间

	// 执行命令
	if _, err := pipe.Exec(); err != nil {
		fmt.Println("pipe.Exec() failed,err:", err)
		return
	}
	fmt.Println("tx_pipeline_counter1's value:", incr.Val())
}

// TxExample2 redis 事务操作示例2
func TxExample2() {
	var incr *redis.IntCmd

	// 使用Pipelined将命令都封装到一个函数中,声省掉了 Exec()步骤
	_, err := rdb.RDB.TxPipelined(func(p redis.Pipeliner) error {
		incr = p.Incr("tx_pipeline_counter2")       // 将key为pipeline_counter的数自增1
		p.Expire("tx_pipeline_counter2", time.Hour) // 设置过期时间
		return nil
	})
	if err != nil {
		fmt.Println("rdb.RDB.Pipelined() failed,err:", err)
		return
	}
	fmt.Println("tx_pipeline_counter2's value:", incr.Val())
}
func main() {
	if err := rdb.InitRedis(); err != nil {
		fmt.Println("rdb.InitRedis failed, err:", err)
		return
	}
	defer rdb.RDB.Close()
	TxExample1()
	TxExample2()
}
