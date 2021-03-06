// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/controller"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/dao"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/service"
	"hb.study/go-lib/di-lib/wire/code/advanced/pkg/database"
	"hb.study/go-lib/di-lib/wire/code/advanced/setting"
)

// Injectors from wire.go:

func initWebApp() (*WebApp, error) {
	engine, err := InitEngine()
	if err != nil {
		return nil, err
	}
	config, err := setting.InitConfig()
	if err != nil {
		return nil, err
	}
	db, err := database.InitMySQL(config)
	if err != nil {
		return nil, err
	}
	userDao := &dao.UserDao{
		DB: db,
	}
	userService := &service.UserService{
		Dao: userDao,
	}
	userController, err := controller.NewUseController(engine, userService)
	if err != nil {
		return nil, err
	}
	webApp := &WebApp{
		Engine: engine,
		config: config,
		user:   userController,
	}
	return webApp, nil
}
