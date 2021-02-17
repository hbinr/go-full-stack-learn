# Go 中 redis set|get|zset的基本使用

[源码](https://github.com/hbinr/go-full-stack-learn/tree/master/redis/code/basic_use)

### set/get 示例

```go
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
```

### zset 示例

**数据结构：**

zset 是有序的，不运行重复的，带有分值 score 的

**数据结构对比：**
![数据结构对比](https://img2018.cnblogs.com/blog/1750617/201910/1750617-20191030142656568-728504027.png)

**Go zset 示例：**

```go
// ZetExample redis zset操作示例
func ZetExample() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
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
```

参考：

[www.liwenzhou.com](https://www.liwenzhou.com/posts/Go/go_redis/#autoid-2-3-0)

[redis 数据类型--zset](https://www.cnblogs.com/sunxuesong/p/11765052.html)

[Redis 为什么这么快？ Redis 的有序集合 zset 的底层实现原理是什么?](https://blog.csdn.net/universsky2015/article/details/102728114?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.edu_weight&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.edu_weight)
