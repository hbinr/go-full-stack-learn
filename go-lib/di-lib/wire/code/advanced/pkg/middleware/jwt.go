package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func JWT(c *gin.Context) {
	// TODO:自定义实现中间件
	fmt.Println("jwt....")
}
