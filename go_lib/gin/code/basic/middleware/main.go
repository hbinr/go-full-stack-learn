package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func fun1(c *gin.Context) {
	fmt.Println("fun1")
}
func fun2(c *gin.Context) {
	fmt.Println("fun2 before")
	c.Next()
	fmt.Println("fun2 after")

}
func fun3(c *gin.Context) {
	c.Abort()
	fmt.Println("fun3")
}
func main() {

	c := gin.Default()
	// 中间件注册方式一：c.Use()
	c.Use(fun1)

	c.GET("/get", func(c *gin.Context) {
		c.String(200, "hello get!")
	}, fun2, fun3) // 另一种注册方式

	c.Run(":8082")

}
