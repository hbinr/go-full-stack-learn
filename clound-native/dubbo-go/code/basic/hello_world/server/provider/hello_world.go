package provider

import (
	"errors"
)

type UserProvider struct {
}

type UserRequest struct {
	ID  int64 `json:"id"`
	Age uint  `json:"age"`
}

type UserReply struct {
	Name string `json:"name"`
}

//	GetUserName GET请求 route:/user/name?id=1
func (h *UserProvider) GetUserName(id string) (res *UserReply, err error) {
	//gxlog.CInfo("id------------:%s", id)
	res = new(UserReply)
	if id != "1" {
		return nil, errors.New("invalid param")
	}
	if id == "1" {
		res.Name = "zhangshan"
	}
	return res, nil
}

//	GetUser POST请求， route:/user/body  PostMan自己拼接json请求，eg:{"id":1,"Age":102}
func (h *UserProvider) GetUser(req *UserRequest) (res *UserReply, err error) {
	//gxlog.CInfo("req:%v", req)
	//user := req[0].(UserRequest)
	res = new(UserReply)
	if req.ID != 1 {
		return nil, errors.New("invalid user")
	}
	if req.Age < 18 {
		return nil, errors.New("未成年") // 底层判断：只要err!=nill，那么就会返回err的内容，此处写 return res 不起作用
	}
	res.Name = "zhangshan"
	return res, nil
}

// Reference 必须实现RPCService接口，才能config.SetProviderService
func (u *UserProvider) Reference() string {
	return "UserProvider"
}
