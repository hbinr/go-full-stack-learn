package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/conf"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const defaultConfigFile = "./config-test.yaml"

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
	return &conf, nil
}
