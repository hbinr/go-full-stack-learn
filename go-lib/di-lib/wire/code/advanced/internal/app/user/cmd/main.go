package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/conf"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/controller"
)

func main() {
	webApp, err := initWebApp()
	if err != nil {
		panic(err)
	}
	webApp.Start()
}

// WebApp represent a web application
type WebApp struct {
	*gin.Engine
	config *conf.Config
	user   *controller.UserController
}

func InitEngine() (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) // 设置公共中间件，可以自定义，举例所以使用了gin自带的
	r.Group("/api")
	return r, nil
}

// Start the web app
func (e *WebApp) Start() {
	e.Run(fmt.Sprintf(":%d", e.config.System.Port))
}
