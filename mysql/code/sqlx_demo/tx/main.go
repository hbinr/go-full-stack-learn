package main

import (
	"database/sql"
	"errors"
	"fmt"

	"hb.study/mysql/code/sqlx_demo/db"
)

// transactionDemo 使用sqlx进行事务操作
func transactionDemo() (err error) {
	var (
		tx *sql.Tx
		rs sql.Result
		n  int64
	)
	if tx, err = db.DB.Begin(); err != nil {
		fmt.Println("db.DB.Begin() failed, err:", err)
		return err
	}
	// defer 来进行事务回滚
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update demo set age=30 where id=?"
	rs, err = tx.Exec(sqlStr1, 1)
	if err != nil {
		fmt.Println("tx.Exec sqlStr1 failed,err:", err)
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		fmt.Println("rs.RowsAffected sqlStr1 failed,err:", err)
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := "Update demo set age=50 where id=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		fmt.Println("tx.Exec sqlStr2 failed,err:", err)
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		fmt.Println("rs.RowsAffected RowsAffected failed,err:", err)
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}

func main() {
	if err := db.InitMysql(); err != nil {
		fmt.Println("sqlx.InitMysql failed,err:", err)
		return
	}
	defer db.DB.Close()
	transactionDemo()
}
