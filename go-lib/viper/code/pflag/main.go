package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

func main() {
	// 设置命令行参数，eg: -c config/config.yaml
	pflag.StringP("configFile", "c", "", "choose config file")
	// 解析命令参数
	pflag.Parse()

	v := viper.New()
	// 绑定flag集到viper中
	v.BindPFlags(pflag.CommandLine)

	// 从viper而不是从pflag检索值
	configFile := v.GetString("configFile")
	if configFile == "" {
		configFile = defaultConfigFile
	}

	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("read config file success, result:", v.GetString("version"))
}
