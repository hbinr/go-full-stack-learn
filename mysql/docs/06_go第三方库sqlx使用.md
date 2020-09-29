# sqlx 库使用

在项目中我们通常可能会使用 database/sql 连接 MySQL 数据库。本文借助使用 sqlx 实现批量插入数据的例子，介绍了 sqlx 中可能被你忽视了的 `sqlx.In` 和 `DB.NamedExec` 方法。

在项目中我们通常可能会使用 database/sql 连接 MySQL 数据库。sqlx 可以认为是 Go 语言内置 database/sql 的超集，它在优秀的内置 database/sql 基础上提供了一组扩展。这些扩展中除了大家常用来查询的

- `Get(dest interface{}, ...) error`
- `Select(dest interface{}, ...) error`

外还有很多其他强大的功能。

## 安装 sqlx

```go
go get github.com/jmoiron/sqlx
```

## 连接数据库

记得匿名导入 mysql 驱动 `_ "github.com/go-sql-driver/mysql"`

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// initMysql 初始化MySQL
func initMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用 :=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = DB.Ping()
	if err != nil {
		return err
	}
	// 设置MySQL相关配置
	DB.SetConnMaxLifetime(20)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)

	return
}

func main() {

	if err := initMysql(); err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	defer DB.Close()

	fmt.Println("mysql connect success!")
}
```

## 查询

### 查询单行数据示例代码如下：

[详细代码](../code/sqlx_demo/crud/main.go)

```go
// Demo Demo结构体
type Demo struct {
	ID   int
	Age  int
	Name string
}

// queryRowDemo 查询单条数据示例
func queryRowDemo() {
	sqlStr := `select id,age,name from demo where id = ?`
	var d Demo
	// Get()返回单条数据
	if err := db.DB.Get(&d, sqlStr, 1); err != nil {
		fmt.Println("db.DB.Get failed, err:", err)
		return
	}
	fmt.Println("queryRowDemo success, data:", d)
}

```

### 查询多行数据示例代码如下：

```go
// queryMultiRowDemo 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := `select id,age,name from demo where id > ?`
	var ds []Demo
	// Select()返回多条数据
	if err := db.DB.Select(&ds, sqlStr, 1); err != nil {
		fmt.Println("db.DB.Select failed, err:", err)
		return
    }

	fmt.Println("queryMultiRowDemo success, data:", ds)
}
```

## 增、删、改

sqlx 中的 exec 方法与原生 sql 中的 exec 使用基本一致

```go
// insertRowDemo 新增一条数据
func insertRowDemo() {
	sqlStr := `insert into demo(age,name) values(?,?)`
	res, err := db.DB.Exec(sqlStr, 20, "Bob")
	if err != nil {
		fmt.Println("insert demo's data failed, err:", err)
		return
	}
	lstID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("res.LastInsertId() failed, err:", err)
		return
	}
	fmt.Printf("insert demo's data success,lstID:%d", lstID)
}

// updateRowDemo 更行数据
func updateRowDemo() {
	sqlStr := `update demo set age = ? where id = ?`
	res, err := db.DB.Exec(sqlStr, 25, 3)
	if err != nil {
		fmt.Println("update demo's data failed, err:", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("update res.RowsAffected error, err:", err)
		return
	}
	fmt.Println("update demo's data success,rowsAffected:", n)
}

// deleteRowDemo 删除数据
func deleteRowDemo() {
	sqlStr := `delete from demo where id = ?`
	res, err := db.DB.Exec(sqlStr, 4)
	if err != nil {
		fmt.Println("delete demo's data failed,err:", err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("delete res.RowsAffected error, err:", err)
		return
	}
	fmt.Println("delete demo's data success,rowsAffected:", n)

}
```

## NamedExec

参考：

https://www.liwenzhou.com/posts/Go/sqlx/
