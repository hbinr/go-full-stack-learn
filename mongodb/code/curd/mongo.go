package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/qiniu/qmgo"
)

type User struct {
	UserID   int64  `bson:"userID"`
	UserName string `bson:"userName"`
}

func main() {
	var (
		err error
		cli *qmgo.QmgoClient
	)
	ctx := context.Background()
	// 1.初始化连接
	cli, err = qmgo.Open(ctx, &qmgo.Config{
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
	// 新增数据
	Insert(ctx, cli)
	// 查询数据
	FindOne(ctx, cli)
	// 修改数据
	UpdateOne(ctx, cli)
	// 删除数据
	//DeleteOne(ctx, cli)
	//DeleteAll(ctx, cli)

}

func Insert(ctx context.Context, cli *qmgo.QmgoClient) error {
	insert := User{
		UserID:   1,
		UserName: "test",
	}
	inserRes, err := cli.InsertOne(ctx, insert)
	if err != nil {
		fmt.Println("新增数据异常，err:", err)
		return err
	}
	fmt.Println("新增数据成功，数据为：", inserRes)

	return nil
}

func FindOne(ctx context.Context, cli *qmgo.QmgoClient) error {
	query := User{}
	err := cli.Find(ctx, bson.M{"userID": 1}).One(&query)
	if err != nil {
		fmt.Println("查询数据异常，err:", err)
		return err
	}
	fmt.Println("查询数据成功，数据为：", query)
	return nil
}

func UpdateOne(ctx context.Context, cli *qmgo.QmgoClient) error {
	err := cli.UpdateOne(ctx, bson.M{"userID": 1}, bson.M{"$set": bson.M{"userName": "test2"}})
	FindOne(ctx, cli)
	return err
}

func DeleteOne(ctx context.Context, cli *qmgo.QmgoClient) error {
	err := cli.Remove(ctx, bson.M{"userID": 1}) // 底层调用DeleteOne
	if err != nil {
		fmt.Println("删除数据异常，err:", err)
		return err
	}
	return nil
}

func DeleteAll(ctx context.Context, cli *qmgo.QmgoClient) error {
	res, err := cli.DeleteAll(ctx, bson.M{"userID": 1}) // 底层调用DeleteMany。返回删除个数
	if err != nil {
		fmt.Println("删除所有数据异常，err:", err)
	}
	fmt.Println("删除的数据个数：", res)
	return nil
}
