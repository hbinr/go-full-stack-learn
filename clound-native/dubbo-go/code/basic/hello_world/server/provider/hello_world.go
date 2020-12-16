package provider

import (
	"errors"

	gxlog "github.com/dubbogo/gost/log"
)

type UserProvider struct {
}

type UserRequest struct {
	ID int64 `json:"id"`
}

type UserReply struct {
	Name string `json:"name"`
}

//	GetUserName GET请求 route:/path?id=12456
func (h *UserProvider) GetUserName(id string) (res *UserReply, err error) {
	gxlog.CInfo("id:%s", id)
	res = new(UserReply)
	if id == "" {
		res.Name = "hello world"
		return res, errors.New("invalid param")
	}
	if id == "1" {
		res.Name = "zhangshan"
	}
	return res, nil
}

//	GetUser POST请求， route:/path?id=12456
func (h *UserProvider) GetUser(req interface{}) (res *UserReply, err error) {
	gxlog.CInfo("req:%v", req)
	user := req.(UserRequest)
	res = new(UserReply)
	if user.ID != 1 {
		res.Name = "test"
		return res, errors.New("invalid user")
	}
	res.Name = "zhangshan"
	return res, nil
}

// Reference 必须实现RPCService接口，才能config.SetProviderService
func (u *UserProvider) Reference() string {
	return "UserProvider"
}
