package dbresolver

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

func TestResolver(t *testing.T) {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	var user User
	if err = db.Where("id = ?", 4).Find(&user).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("user: ", user)
}

const db1_dsn = "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
const db2_dsn = "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
const db3_dsn = "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
const db4_dsn = "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
const db5_dsn = "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
const db6_dsn = "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
const db7_dsn = "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"
const db8_dsn = "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True"

type User struct {
	Id       int
	UserName string
}

func (u User) TableName() string {
	return "user"
}

type Address struct{}
type Order struct{}
type Product struct{}

func initDB() (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(db1_dsn), &gorm.Config{
		// 开启SQL日志: logger.Info
		// 不开启SQL日志: logger.Silent
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加复数形式，false默认加
		},
	})

	db.Use(dbresolver.Register(dbresolver.Config{
		// `db2` 作为 sources，`db3`、`db4` 作为 replicas
		Sources:  []gorm.Dialector{mysql.Open(db2_dsn)},
		Replicas: []gorm.Dialector{mysql.Open(db3_dsn), mysql.Open(db4_dsn)},
		// sources/replicas 负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}).Register(dbresolver.Config{
		// `db1` 作为 sources（DB 的默认连接），对于 `User`、`Address` 使用 `db5` 作为 replicas
		Replicas: []gorm.Dialector{mysql.Open(db5_dsn)},
	}, &User{}, &Address{}).Register(dbresolver.Config{
		// `db6`、`db7` 作为 sources，对于 `orders`、`Product` 使用 `db8` 作为 replicas
		Sources:  []gorm.Dialector{mysql.Open(db6_dsn), mysql.Open(db7_dsn)},
		Replicas: []gorm.Dialector{mysql.Open(db8_dsn)},
	}, &Order{}, &Product{}, "secondary"))

	return db, err
}
