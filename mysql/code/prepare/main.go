package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type Demo struct {
	ID   int
	Age  int
	Name string
}

func initMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&parseTime=True"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

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

	sqlStr := "select * from demo where id > ? "
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("DB.Prepare failed, err:", err)
		return
	}
	// 非常重要：释放命令部分资源，当MySQL连接数被占满，资源不合理释放
	// 会导致sql一直等待可利用资源却资源沾满而无法执行
	defer stmt.Close()

	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Println("stmt.Query failed, err:", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var d Demo
		err := rows.Scan(&d.ID, &d.Age, &d.Name)
		if err != nil {
			fmt.Println("rows.Scan failed,err:", err)
			return
		}
		fmt.Println("Demo's data:", d)
	}
}
