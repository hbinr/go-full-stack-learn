package rdb

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

// InitFailoverRedis 连接 Redis 哨兵模式
func InitFailoverRedis() (err error) {
	RDB = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

var RDBC *redis.ClusterClient

// InitClusterRedis 连接 Redis 集群
func InitClusterRedis() (err error) {
	RDBC = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
