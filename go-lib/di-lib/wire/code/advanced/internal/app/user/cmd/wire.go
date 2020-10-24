//+build wireinject

package main

import (
	"github.com/google/wire"
	"hb.study/go-lib/di-lib/wire/code/advanced/pkg/database"
	"hb.study/go-lib/di-lib/wire/code/advanced/setting"
	"hb.study/go-lib/di-lib/wire/code/advanced/user"
	"hb.study/go-lib/di-lib/wire/code/advanced/user/controller"
)

// initWebApp 注入函数，自定义的函数直接注入就行，不需要使用wire set
func initWebApp() (*WebApp, error) {
	// 逻辑顺序入参，未用到的依赖不需要注入
	wire.Build(
		InitEngine,                    // 初始化web引擎，自定义
		setting.InitConfig,            // 加载日志，自定义
		database.InitMySQL,            // 获取gorm.DB，自定义
		user.Set,                      // user业务provider，wire生成
		controller.NewUseController,   // 获取业务controller，自定义
		wire.Struct(new(WebApp), "*"), // WebApp provider，wire生成
	)
	// 返回值不用管。直接返回nil就行
	return nil, nil
}
