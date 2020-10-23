package controller

import (
	"fmt"
	"strconv"

	"github.com/gogf/gf/util/gconv"

	"hb.study/go-lib/di-lib/wire/code/advanced/user/model"

	"hb.study/go-lib/di-lib/wire/code/advanced/pkg/ginx"
	"hb.study/go-lib/di-lib/wire/code/advanced/user/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	engine      *gin.Engine
	userService service.UserServiceProvider
}

// SignUp 新增用户
func (u *UserController) SignUp(c *gin.Context) {
	var (
		err    error
		uParam model.SignUpInput
		uModel service.UserDto
	)

	_ = c.ShouldBind(&uParam)
	if err = gconv.Struct(uParam, &uModel); err != nil {
		ginx.FailWithMessage("数据转换异常", c)
		return
	}
	if err = u.userService.Insert(&uModel); err != nil {
		ginx.FailWithMessage("注册用户异常", c)
		return
	}
	ginx.OkDetailed(&uModel, "注册用户成功", c)
}

// Get 根据id获取用户
func (u *UserController) Get(c *gin.Context) {
	var (
		id   int
		err  error
		user *service.UserDto
	)
	idStr := c.Query("id")
	if id, err = strconv.Atoi(idStr); err != nil {
		ginx.FailWithMessage("服务器异常", c)
		fmt.Println("UserController.Get strconv.Atoi failed, err:", err)
		return
	}
	if user, err = u.userService.SelectById(int64(id)); err != nil {
		ginx.FailWithMessage("获取用户失败", c)
		return
	}
	ginx.OkDetailed(&user, "获取用户成功", c)
}
