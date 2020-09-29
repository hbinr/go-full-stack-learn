package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB 数据库句柄
var DB *sqlx.DB

// InitMysql 初始化MySQL连接
func InitMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&parseTime=True"
	// sqlx.Connect() 底层做了Open和Ping
	if DB, err = sqlx.Connect("mysql", dsn); err != nil {
		fmt.Println("sqlx.Connect failed, err:", err)
		return
	}
	// 设置MySQL相关配置
	DB.SetConnMaxLifetime(20)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
	fmt.Println("mysql connect success......")
	return
}
