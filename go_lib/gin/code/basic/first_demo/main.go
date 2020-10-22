package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/student", func(c *gin.Context) {
		c.String(http.StatusOK, "使用String()做出响应:%s", "GET 请求..")
	})

	r.POST("/student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "add request success",
		})
	})

	r.PUT("/student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "update request success",
		})
	})

	r.DELETE("/student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete request success",
		})
	})

	r.Run()
}
