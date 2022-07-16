package main

import (
	"fmt"

	"hb.study/mysql/code/sqlx_demo/db"
)

// Demo Demo结构体
type Demo struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
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
	// SlectByName()返回多条数据
	if err := db.DB.Select(&ds, sqlStr, 1); err != nil {
		fmt.Println("db.DB.SlectByName failed, err:", err)
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

// insertRowByNameExecDemo 使用NamedExec()操作sql，针对增删改
func insertRowByNameExecDemo() {
	sqlStr := `insert into demo(age,name) values(:age,:name)`
	// 使用结构体
	ronger := Demo{Age: 19, Name: "ronger"}

	res, err := db.DB.NamedExec(sqlStr, &ronger)
	if err != nil {
		fmt.Println("insert demo's data failed,err:", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("res.LastInsertId() failed, err:", err)
		return
	}
	fmt.Println("insert demo's data success, 5id:", id)

}

// queryRowByNamedQueryDemo 使用NamedExec()操作sql，针对查
func queryRowByNamedQueryDemo() {
	sqlStr := `SELECT * from demo WHERE name=:name`
	demo := Demo{
		Name: "Bob",
	}
	rows, err := db.DB.NamedQuery(sqlStr, &demo)
	if err != nil {
		fmt.Println("select demo's data by NamedQuery failed,err:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		// 方式一：使用StructScan，一次扫描一行
		var d Demo
		err := rows.StructScan(&d)
		if err != nil {
			fmt.Println("rows.StructScan failed, err:", err)
			continue
		}
		fmt.Println("row's data :", d)

		// 方式二：使用SliceScan，一次扫描多行
		res, err := rows.SliceScan()
		if err != nil {
			fmt.Println("rows.SliceScan failed,err:", err)
			return
		}
		fmt.Println("rows's res :", res)
	}
}

func main() {
	if err := db.InitMysql(); err != nil {
		fmt.Println("sqlx.InitMysql failed,err:", err)
		return
	}
	defer db.DB.Close()
	queryRowDemo()
	queryMultiRowDemo()
	// insertRowDemo()
	// updateRowDemo()
	// deleteRowDemo()
	// insertRowByNameExecDemo()
	queryRowByNamedQueryDemo()
}
