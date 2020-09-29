package main

import (
	"fmt"
	"go-full-stack-learn/mysql/code/sqlx_demo/db"
)

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
	fmt.Println("insert demo's data success,lstID", lstID)
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

func main() {
	if err := db.InitMysql(); err != nil {
		fmt.Println("sqlx.InitMysql failed,err:", err)
		return
	}
	defer db.DB.Close()
	queryRowDemo()
	queryMultiRowDemo()
	insertRowDemo()
	updateRowDemo()
	deleteRowDemo()
}
