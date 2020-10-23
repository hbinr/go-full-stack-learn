package setting

import (
	"fmt"

	"github.com/spf13/viper"
	"hb.study/go-lib/di-lib/wire/code/advanced/user/conf"
)

func InitConfig() (*conf.Config, error) {
	var (
		conf conf.Config
		err  error
	)
	viper.SetConfigFile("config-test.yaml")
	viper.AddConfigPath(".") // 在cd到 cmd目录下，执行main.go
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println("viper.Unmarshal failed,err", err)
		return nil, err
	}
	return &conf, nil
}
