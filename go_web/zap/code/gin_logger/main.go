package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.New()

	logger, _ := zap.NewProduction()

	// 添加一个ginzap中间件，该中间件：
	//   - 记录所有请求，例如合并的访问和错误日​​志。
	//   - 以stdout格式输出.
	//   - 具有UTC时间格式的RFC3339.
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// recover掉项目可能出现的panic，并使用zap记录相关日志
	//   - 第二个参数表示：是否输出堆栈信息，true则输出
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	r.Run(":8080")
}
