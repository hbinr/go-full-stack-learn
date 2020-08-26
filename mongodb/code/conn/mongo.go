package conn

import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
)

func InitMongo() {
	ctx := context.Background()
	// 1.初始化连接
	cli, err := qmgo.Open(ctx, &qmgo.Config{
		Uri:      "mongodb://localhost:27017",
		Database: "cron",
		Coll:     "user"})
	if err != nil {
		fmt.Println("连接MongoDB异常，err:", err)
		return
	}
	// 2.在初始化成功后，defer 来关闭连接
	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
}
