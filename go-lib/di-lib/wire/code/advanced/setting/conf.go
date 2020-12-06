package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/conf"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const defaultConfigFile = "./config-httprouter_test.yaml"

// InitConfig init config
func InitConfig() (*conf.Config, error) {
	var (
		conf conf.Config
		err  error
	)
	pflag.StringP("conf", "c", "", "choose conf file.")
	pflag.Parse()

	// 优先级: 命令行 > 环境变量 > 默认值
	v := viper.New()
	v.BindPFlags(pflag.CommandLine)

	configFile := v.GetString("conf")
	if configFile == "" {
		configFile = defaultConfigFile
	}

	v.SetConfigFile(configFile)
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error conf file: %s", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("conf file changed:", e.Name)
		if err := v.Unmarshal(&conf); err != nil {
			fmt.Println("v.OnConfigChange -> v.Unmarshal(&conf) failed,err:", err)
		}
	})
	if err := v.Unmarshal(&conf); err != nil {
		fmt.Println("v.Unmarshal(&conf) failed,err:", err)
	}

	fmt.Println("Config sets success，Conf:", &conf)
	//
	//viper.SetConfigFile("config-httprouter_test.yaml")
	//viper.AddConfigPath(".") // 在cd到 cmd目录下，执行main.go
	//if err = viper.ReadInConfig(); err != nil {
	//	panic(fmt.Errorf("Fatal error config file: %s", err))
	//}
	//if err := viper.Unmarshal(&conf); err != nil {
	//	fmt.Println("viper.Unmarshal failed,err", err)
	//	return nil, err
	//}
	return &conf, nil
}
