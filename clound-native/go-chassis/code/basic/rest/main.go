package main

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"

	"hb.study/clound-native/go-chassis/code/basic/rest/server"
)

func main() {
	chassis.RegisterSchema("rest", &server.UserController{}) // 注册对外暴露的协议+服务
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
