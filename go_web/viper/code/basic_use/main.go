package main

import (
	"fmt"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// ReadConfigDemo viper读取配置文件示例
func ReadConfigDemo() {
	// 设置配置文件名
	// viper.SetConfigName("config")

	// 设置查找配置文件所在的路径，可以配置多个,按代码顺序寻找
	viper.AddConfigPath(".")

	// 设置配置文件类型
	// viper.SetConfigType("yaml")

	// 设置配置文件名，使用方法后，便不需要配置 SetConfigName和SetConfigType
	viper.SetConfigFile("config.yaml")

	// 查找并读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			fmt.Println("未找到配置文件，err:", err)
			return
		}
		// 配置文件被找到，但产生了另外的错误
		fmt.Println("配置文件被找到，但产生了另外的错误，err:", err)
		return
	}

	// 配置文件找到并成功解析
	fmt.Println("获取配置信息成功，version:", viper.GetString("version"))
}

// WriteConfigDemo viper写入配置文件示例
func WriteConfigDemo() {
	viper.SetConfigFile("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.WriteConfig(); err != nil {
		fmt.Println("viper SafeWriteConfig failed,err:", err)
		return
	}
	fmt.Println("viper write config success, version's value:", viper.GetString("version"))
}

// WatchConfigDemo viper监控配置文件示例
func WatchConfigDemo() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	

	// 实时监控配置
	viper.WatchConfig()
	// 配置文件发生变更之后会调用回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// 验证监控配置示例
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})
	r.Run()
}	return
	}

	// 实时监控配置
	viper.WatchConfig()
	// 配置文件发生变更之后会调用回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// 验证监控配置示例
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})
	r.Run()
}

func main() {
	// ReadConfigDemo()
	// WriteConfigDemo()
	WatchConfigDemo()

}
