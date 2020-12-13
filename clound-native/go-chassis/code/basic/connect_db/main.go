package main

import (
	"fmt"

	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/config/model"
	"github.com/go-chassis/go-chassis/v2/storage"
	"github.com/go-chassis/openlog"

	"hb.study/clound-native/go-chassis/code/basic/rest/server"
)

func main() {
	chassis.RegisterSchema("rest", &server.UserController{}) // 注册对外暴露的协议+服务
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}

	fmt.Println("redis--:", archaius.GetString("servicecomb.redis.link", "1"))
	fmt.Printf("mysql--%v:", NewMysqlConf())

	chassis.Run()
}
func NewMysqlConf() *storage.Options {
	globalDefinition := &model.GlobalCfg{}
	err := archaius.UnmarshalConfig(&globalDefinition)
	if err != nil {
		return nil
	}
	return &globalDefinition.ServiceComb.Options
}
