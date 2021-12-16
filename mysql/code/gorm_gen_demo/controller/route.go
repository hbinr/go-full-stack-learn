package controller

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine, userCtl userController) {
	r.GET("/create", userCtl.CreateUser)
	r.GET("/", userCtl.GetUser)
	r.GET("/condition", userCtl.GetUserByCondition)
}
