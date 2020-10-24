package controller

import (
	"fmt"
	"strconv"

	"hb.study/go-lib/di-lib/wire/code/advanced/pkg/middleware"

	"github.com/gogf/gf/util/gconv"

	"hb.study/go-lib/di-lib/wire/code/advanced/user/model"

	"hb.study/go-lib/di-lib/wire/code/advanced/pkg/ginx"
	"hb.study/go-lib/di-lib/wire/code/advanced/user/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	engine      *gin.Engine
	userService service.IUserService
}

func NewUseController(e *gin.Engine, us service.IUserService) (*UserController, error) {
	user := &UserController{
		engine:      e,
		userService: us,
	}
	g := e.Group("/user")
	g.Use(middleware.JWT) // 设置user私有中间件
	{
		g.POST("/signup", user.SignUp)
		g.GET("/get", user.Get)
		g.GET("/delete", user.Delete)
	}
	return user, nil
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

// Delete 根据id删除用户
func (u *UserController) Delete(c *gin.Context) {
	var (
		id  int
		err error
	)
	idStr := c.Query("id")
	if id, err = strconv.Atoi(idStr); err != nil {
		ginx.FailWithMessage("服务器异常", c)
		return
	}
	if !u.userService.Delete(int64(id)) {
		ginx.FailWithMessage("删除用户失败", c)
		return
	}
	ginx.OkWithMessage("删除用户成功", c)
}
