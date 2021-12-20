package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hb.study/mysql/code/gorm_gen_demo/controller"
	"hb.study/mysql/code/gorm_gen_demo/repository"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("db init error")
	}

	userRepo := repository.NewUserRepo(db)
	userCtl := controller.NewUserController(userRepo)
	r := gin.Default()
	controller.InitRouter(r, userCtl)

	fmt.Println("http://127.0.0.1:8092")
	if err = r.Run(":8092"); err != nil {
		panic(err)
	}
}
