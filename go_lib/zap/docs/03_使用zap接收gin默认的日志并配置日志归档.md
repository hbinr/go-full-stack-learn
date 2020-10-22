# 使用zap接收gin框架默认的日志并配置日志归档

本文介绍了在基于gin框架开发的项目中如何配置并使用zap来接收并记录gin框架默认的日志和如何配置日志归档。

我们在基于gin框架开发项目时通常都会选择使用专业的日志库来记录项目中的日志，go语言常用的日志库有zap、logrus等。

但是我们该如何在日志中记录gin框架本身输出的那些日志呢？

## gin默认的中间件
首先我们来看一个最简单的gin项目：

```go
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String("hello q1mi!")
	})
	r.Run()
}
```

接下来我们看一下gin.Default()的源码：
```go
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
```
也就是我们在使用`gin.Default()`的同时是用到了gin框架内的两个默认中间件`Logger()`和`Recovery()`。

其中`Logger()`是把gin框架本身的日志输出到标准输出（我们本地开发调试时在终端输出的那些日志就是它的功劳），而`Recovery()`是在程序出现panic的时候恢复现场并写入500响应的。

## 基于zap的中间件

我们可以使用github上有别人封装好的https://github.com/gin-contrib/zap。

```go
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

	// 示例1 
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})
	// 示例2
	r.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	r.Run(":8080")
}
```