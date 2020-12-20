package provider

import (
	"errors"

	gxlog "github.com/dubbogo/gost/log"
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

// GetUserName GET请求 route:/user/name?id=1
func (h *UserProvider) GetUserName(id string) (res *UserReply, err error) {
	// gost v1.9.2
	gxlog.CInfo("id------------:%s", id)
	res = new(UserReply)
	if id != "1" {
		res.Name = "犯错了"
		return res, errors.New("invalid param")
	}
	if id == "1" {
		res.Name = "zhangshan"
	}
	return res, nil
}

// GetUser POST请求， route:/user/body  PostMan自己拼接json请求，eg:{"id":1,"Age":102}
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

// UpdateUser POST请求， route:/user/update  PostMan自己拼接json请求，eg:{"id":1,"Age":102}
func (h *UserProvider) UpdateUser(req []interface{}) (res *UserReply, err error) {
	gxlog.CInfo("UpdateUser req:%v", req)
	user := req[0].(UserRequest)
	// req interface{}或req []interface{}都会报以下错误
	// interface conversion: interface {} is map[string]interface {}, not provider.UserRequest

	res = new(UserReply)
	if user.ID != 1 {
		return nil, errors.New("invalid user")
	}
	if user.Age < 18 {
		return nil, errors.New("未成年") // 底层判断：只要err!=nill，那么就会返回err的内容，此处写 return res 不起作用
	}
	res.Name = "zhangshan"
	return res, nil
}

// UpdateUser2 POST请求， route:/user/update  PostMan自己拼接json请求，eg:{"id":1,"Age":102}
func (h *UserProvider) UpdateUser2(req interface{}) (res *UserReply, err error) {
	gxlog.CInfo("UpdateUser req:%v", req)

	// 底层返回的是：map[string]interface{}，那就断言一下它
	user := req.(map[string]interface{})
	gxlog.CInfo("UpdateUser t:%v", user)
	res = new(UserReply)
	if user["id"] != 1 {
		return nil, errors.New("invalid user")
	}
	// if user["age"] < 18 { // < > 这类比较大小炒作不能用于interface{}，因此入参的定义很关键，毕竟断言能完全提供空接口性能
	// 	return nil, errors.New("未成年") // 底层判断：只要err!=nill，那么就会返回err的内容，此处写 return res 不起作用
	// }
	res.Name = "zhangshan"
	return res, nil
}

// Reference 必须实现RPCService接口，才能config.SetProviderService
func (u *UserProvider) Reference() string {
	return "UserProvider"
}
