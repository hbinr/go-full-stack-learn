package main

import (
	"database/sql/driver"
	"fmt"
	"go-full-stack-learn/mysql/code/sqlx_demo/db"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Demo Demo结构体
type Demo struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

// Value 使用sqlx.In 必须实现的接口
func (d Demo) Value() (driver.Value, error) {
	return []interface{}{d.Age, d.Name}, nil
}

// BatchInsertDemos sqlx.In 批量新增数据
func BatchInsertDemos(ds []interface{}) error {
	sqlStr := `INSERT  INTO  demo(age,name) values(?),(?),(?)` // 多个一个占位符
	query, args, err := sqlx.In(sqlStr, ds...)                 // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开ds
	if err != nil {
		fmt.Println("sqlx.In insert demo's data failed,err:", err)
		return err
	}
	fmt.Println("查看生成的querystring:", query)
	fmt.Println("查看生成的args", args)

	_, err = db.DB.Exec(query, args...)
	if err != nil {
		fmt.Println("db.DB.Exec insert demo's data failed,err:", err)
		return err
	}
	return nil
}

// BatchInsertDemos2 NamedExec批量新增数据，不需要实现 driver.Valuer 接口
func BatchInsertDemos2(demos []*Demo) error {
	_, err := db.DB.NamedExec("INSERT INTO demo (age, name) VALUES (:age, :name)", demos)
	return err
}

// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int) (demos []Demo, err error) {
	query, args, err := sqlx.In(`SELECT * FROM demo WHERE id IN (?)`, ids)
	if err != nil {
		fmt.Println("QueryByIDs select demo's data failed,err:", err)
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.DB.Rebind(query)
	fmt.Println("查询的sql为：", query)
	fmt.Println("拼接的参数为：", args)
	// Select() 查询多条数据
	err = db.DB.Select(&demos, query, args...)
	return
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (demos []Demo, err error) {
	// 将int型切片转成string切片参数，用作sql参数
	strIDs := make([]string, 0, len(ids))
	for _, v := range ids {
		strIDs = append(strIDs, strconv.Itoa(v))
	}

	// 拼接sql，一定要加strIDs用 ','隔开
	query, args, err := sqlx.In(`SELECT * FROM demo WHERE id IN (?) ORDER BY FIND_IN_SET(id,?)`,
		ids, strings.Join(strIDs, ","))
	if err != nil {
		fmt.Println("QueryAndOrderByIDs select demo's data failed,err:", err)
		return
	}
	query = db.DB.Rebind(query)
	err = db.DB.Select(&demos, query, args...)
	return
}

func main() {
	if err := db.InitMysql(); err != nil {
		fmt.Println("sqlx.InitMysql failed,err:", err)
		return
	}
	defer db.DB.Close()
	d1 := Demo{Age: 18, Name: "小明"}
	d2 := Demo{Age: 28, Name: "小明2"}
	d3 := Demo{Age: 38, Name: "小明3"}
	demos1 := []interface{}{&d1, &d2, &d3}
	BatchInsertDemos(demos1)

	demos2 := []*Demo{&d1, &d2, &d3}
	BatchInsertDemos2(demos2)

	res, err := QueryByIDs([]int{1, 2, 3, 5})
	if err != nil {
		fmt.Println("QueryByIDs failed，err:", err)
		return
	}
	fmt.Println("QueryByIDs success,data:", res)

	res, err = QueryAndOrderByIDs([]int{1, 6, 7, 3, 5})
	if err != nil {
		fmt.Println("QueryAndOrderByIDs failed，err:", err)
		return
	}
	fmt.Println("QueryAndOrderByIDs success,data:", res)

}
