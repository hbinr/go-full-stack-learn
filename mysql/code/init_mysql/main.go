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
