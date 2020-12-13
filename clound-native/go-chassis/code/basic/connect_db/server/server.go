package server

import (
	"net/http"

	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
)

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UserController struct {
}

// PostUser .
func (r *UserController) PostUser(ctx *restful.Context) {
	var (
		req LoginRequest
		err error
	)
	if err = ctx.ReadEntity(&req); err != nil {
		openlog.Error("server:LoginRequest ReadEntity failed" + err.Error())
		return
	}

	ctx.WriteJSON(&req, "application/json") // 既然已经是WriteJson了，为什么还要指定Content-Type类型？
}

// URLPatterns 帮助响应相应的API调用，必须实现，底层会自动掉接口。类似我们平时写的在RegisterHandle中设置路由
func (u *UserController) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodPost, Path: "/user", ResourceFunc: u.PostUser},
	}
}
