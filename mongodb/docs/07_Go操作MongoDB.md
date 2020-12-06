# Go 操作 MongoDB

GitHub：https://github.com/qiniu/qmgo

## 使用 Qmgo 作为 MongoDB driver

Qmgo 是一款 Go 语言的 MongoDB driver，它基于 MongoDB 官方 driver 开发实现，同时使用更易用的接口设计，比如参考 mgo （比如 mgo 的链式调用）。

- Qmgo 能让用户以更优雅的姿势使用 MongoDB 的新特性。
- Qmgo 是从 mgo 迁移到新 MongoDB driver 的第一选择，对代码的改动影响最小。

### 要求

Go 1.10 及以上。
MongoDB 2.6 及以上。

### 功能

- 文档的增删改查
- 创建链接时支持配置: 连接池、pool Monitor、Auth、ReadPreference
- 索引配置、删除
- Sort、limit、count、select、distinct
- Cursor
- 聚合 Aggregate
- 事务
- 预定义操作符

### 安装

推荐方式是使用 go mod，通过在源码中 import github.com/qiniu/qmgo 来自动安装依赖。

当然，通过下面方式同样可行：

```go
go get github.com/qiniu/qmgo
```

## 开始使用

### import 并新建连接

```go
import(
    "context"

    "github.com/qiniu/qmgo"
)

ctx := context.Background()
client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
db := client.Database("class")
coll := db.Collection("app")
```

如果你的连接是指向固定的 database 和 collection，我们推荐使用下面的更方便的方法初始化连接，后续操作都基于 cli 而不用再关心 database 和 collection

```go
cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "class", Coll: "app"})
```

后面都会基于 cli 来举例，如果你使用第一种传统的方式进行初始化，根据上下文，将 `cli` 替换成 `client`、`db` 或 `coll` 即可

在初始化成功后，请 defer 来关闭连接

```go
defer func() {
    if err = cli.Close(ctx); err != nil {
        panic(err)
    }
}()
```

### 创建索引

**做操作前，我们先初始化一些数据**：

```go
type UserInfo struct {
	Name   string `bson:"name"`
	Age    uint16 `bson:"age"`
	Weight uint32 `bson:"weight"`
}

var oneUserInfo = UserInfo{
	Name:   "xm",
	Age:    7,
	Weight: 40,
}
```

**创建索引**：

```go
cli.EnsureIndexes(ctx, []string{}, []string{"age", "name,weight"})
```

### 插入一个文档

```go
// insert one document
result, err := cli.Insert(ctx, oneUserInfo)
```

返回结果为 `*qmgo.InsertOneResult`，是一个封装了 ObjectID 的结构体，如：`&{ObjectID("5f461df00bc3156bd6a65b7b")}`

### 查找一个文档

```go
// find one document
one := UserInfo{}
err = cli.Find(ctx, bson.M{"name": oneUserInfo.Name}).One(&one)
```

### 删除文档

- 删除单个文档，底层调用 DeleteOne

```go
err = cli.Remove(ctx, bson.M{"age": 7})
```

- 删除多个文档，底层调用 DeleteMany，返回删除个数

```go
res,err := cli.DeleteAll(ctx, bson.M{"age": 7})

fmt.Println(res)
```

打印结果：

> &{3}

### 插入多条数据

```go
// multiple insert
var batchUserInfoI = []interface{}{
	UserInfo{Name: "a1", Age: 6, Weight: 20},
	UserInfo{Name: "b2", Age: 6, Weight: 25},
	UserInfo{Name: "c3", Age: 6, Weight: 30},
	UserInfo{Name: "d4", Age: 6, Weight: 35},
	UserInfo{Name: "a1", Age: 7, Weight: 40},
	UserInfo{Name: "a1", Age: 8, Weight: 45},
}
result, err = cli.Collection.InsertMany(ctx, batchUserInfoI)
```

### 批量查找、Sort 和 Limit

```go
// find all 、sort and limit
batch := []UserInfo{}
cli.Find(ctx, bson.M{"age": 6}).Sort("weight").Limit(7).All(&batch)
```

### 修改文档

```go
// UpdateOne one
err := cli.UpdateOne(ctx, bson.M{"name": "d4"}, bson.M{"\$set": bson.M{"age": 7}})

// UpdateAll
result, err := cli.UpdateAll(ctx, bson.M{"age": 6}, bson.M{"\$set": bson.M{"age": 10}})
```

### 汇总 Count

```go
count, err := cli.Find(ctx, bson.M{"age": 6}).Count()
```

### Select

```go
err := cli.Find(ctx, bson.M{"age": 10}).Select(bson.M{"age": 1}).One(&one)
```

### 聚合 Aggregate

```go
matchStage := bson.D{{"$match", []bson.E{{"weight", bson.D{{"$gt", 30}}}}}}
groupStage := bson.D{{"$group", bson.D{{"_id", "$name"}, {"total", bson.D{{"$sum", "$age"}}}}}}
var showsWithInfo []bson.M
err = cli.Aggregate(context.Background(), Pipeline{matchStage, groupStage}).All(&showsWithInfo)
```

### Pool Monitor

```go
poolMonitor := &event.PoolMonitor{
	Event: func(evt *event.PoolEvent) {
		switch evt.Type {
		case event.GetSucceeded:
			fmt.Println("GetSucceeded")
		case event.ConnectionReturned:
			fmt.Println("ConnectionReturned")
		}
	},
}
cli, err := Open(ctx, &Config{Uri: URI, Database: DATABASE, Coll: COLL, PoolMonitor: poolMonitor})

```

### 事务

简单而强大的事务, 同时还有超时和重试等功能:

```go
callback := func(sessCtx context.Context) (interface{}, error) {
    // 重要：确保事务中的每一个操作，都使用传入的sessCtx参数
    if _, err := cli.InsertOne(sessCtx, bson.D{{"abc", int32(1)}}); err != nil {
        return nil, err
    }
    if _, err := cli.InsertOne(sessCtx, bson.D{{"xyz", int32(999)}}); err != nil {
        return nil, err
    }
    return nil, nil
}
result, err = cli.DoTransaction(ctx, callback)
```

[关于事务的更多内容](https://github.com/qiniu/qmgo/wiki/Transactions)

### 预定义操作符

```go
// aggregate
matchStage := bson.D{{operator.Match, []bson.E{{"weight", bson.D{{operator.Gt, 30}}}}}}
groupStage := bson.D{{operator.Group, bson.D{{"_id", "$name"}, {"total", bson.D{{operator.Sum, "$age"}}}}}}
var showsWithInfo []bson.M
err = cli.Aggregate(context.Background(), Pipeline{matchStage, groupStage}).All(&showsWithInfo)
```
