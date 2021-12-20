package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, userCtl userController) {
	r.GET("/create", userCtl.CreateUser)
	r.GET("/", userCtl.GetUser)
	r.GET("/condition", userCtl.GetUserByCondition)
	r.GET("/select-filed", userCtl.GetUserBySelectFiled)
	r.GET("/page", userCtl.GetUserByPage)
	r.GET("/pluck", userCtl.GetSingleFiledByPluck)
	r.GET("/update", userCtl.UpdateSingleFiled)
	r.GET("/update-select", userCtl.UpdateSelectFiled)
	r.GET("/update-omit", userCtl.UpdateOmitFiled)
	r.GET("/tx", userCtl.DeleteAndUpdate)
	r.GET("/tx-begin", userCtl.DeleteAndCreate)

}
