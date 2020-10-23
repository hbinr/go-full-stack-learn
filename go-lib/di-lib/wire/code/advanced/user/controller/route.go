package controller

import (
	"github.com/gin-gonic/gin"
	"hb.study/go-lib/di-lib/wire/code/advanced/user/service"
)

func NewUseController(e *gin.Engine, us service.UserServiceProvider) (*UserController, error) {
	user := &UserController{
		engine:      e,
		userService: us,
	}
	g := e.Group("/user")
	{
		g.POST("/signup", user.SignUp)
		g.GET("/get", user.Get)
	}
	return user, nil
}
