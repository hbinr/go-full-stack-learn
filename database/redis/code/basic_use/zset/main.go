package main

import (
	"fmt"

	"hb.study/database/redis/code/rdb"

	"github.com/go-redis/redis"
)

// ZetExample redis zset操作示例
func ZetExample() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}

	// ZADD  添加有序集合
	num, err := rdb.RDB.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := rdb.RDB.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Println("ZIncrBy golang's score failed,err:", err)
		return
	}
	fmt.Println("ZIncrBy golang's score success,newScore:", newScore)

	// ZRevRangeWithScores 按score从大到小排序，
	res, err := rdb.RDB.ZRevRangeWithScores(zsetKey, 0, 2).Result() // 0,2 表示取前三名
	if err != nil {
		fmt.Println("ZRevRangeWithScores get the key 'language_rank' failed,err:", err)
		return
	}
	fmt.Println("ZRevRangeWithScores get the key 'language_rank' success,:", res)
	for _, z := range res {
		fmt.Println(z.Member, z.Score)
	}

	// ZRangeByScoreWithScores 命令：zrangebyscore zsetKey 95 100（ 通过score取出95-100的language，从小到大）
	rangeBy := redis.ZRangeBy{ //构造区间结构体
		Min: "95",
		Max: "100",
	}
	res, err = rdb.RDB.ZRangeByScoreWithScores(zsetKey, rangeBy).Result()
	if err != nil {
		fmt.Println("ZRangeByScoreWithScores get the key 'language_rank' which score in [95-100] failed,err:", err)
		return
	}
	fmt.Println("ZRangeByScoreWithScores get the key 'language_rank' which score in [95-100] success,:", res)
	for _, z := range res {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	if err := rdb.InitRedis(); err != nil {
		fmt.Println("rdb.InitRedis failed, err:", err)
		return
	}
	defer rdb.RDB.Close()
	ZetExample()

}
