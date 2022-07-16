# Go 操作 MySQL

## Go 连接 MySQL

### 下载依赖

```go
go get -u github.com/go-sql-driver/mysql
```

### 使用 MySQL 驱动

```go

func Open(driverName, dataSourceName string) (\*DB, error)
```

Open 打开一个 dirverName 指定的数据库，dataSourceName 指定数据源，一般至少包括数据库文件名和其它连接必要的信息。

Open 函数只是验证其参数格式是否正确，实际上并不创建与数据库的连接。如果要检查数据源的名称是否真实有效，应该调用 Ping 方法。

返回的 DB 对象可以安全地被多个 goroutine 并发使用，并且维护其自己的空闲连接池。因此，Open 函数应该仅被调用一次，很少需要关闭这个 DB 对象。

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
		fmt.Printf("setting db failed,err:%v\n", err)
		return
	}
	defer DB.Close()

	fmt.Println("mysql connect success!")
}
```

其中 sql.DB 是表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。它内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个 goroutine 同时使用。

### 数据库相关配置

#### SetMaxOpenConns

设置到数据库的最大打开连接数，如果 MaxIdleConns 大于 0 并且新的 MaxOpenConns 小于 MaxIdleConns，那么将减少 MaxIdleConns 以匹配新的 de MaxOpenConns 限制。

如果 n <= 0，则打开的连接数没有限制。默认值为 0（无限制）。

#### SetConnMaxLifetime

设置可以重用连接的最长时间，过期的连接可能会在重新使用之前延迟关闭。

#### SetMaxIdleConns

设置空闲连接池中的最大连接数，如果 MaxOpenConns 大于 0 但小于新的 MaxIdleConns，那么将减少新的 MaxIdleConns 以匹配 MaxOpenConns 限制。

如果 n <= 0，则不保留空闲连接。当前默认的最大空闲连接数为 2。
