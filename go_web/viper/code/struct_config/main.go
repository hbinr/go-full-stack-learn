package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host   string `mapstructure:"host"`
	DbName string `mapstructure:"dbname"`
	Port   int    `mapstructure:"port"`
}

func main() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("can't find config file, err:", err)
			return
		}
		fmt.Println("viper.ReadInConfig failed, err:", err)
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Println("viper.Unmarshal failed, err:", err)
	}
	fmt.Printf("viper.Unmarshal success, config:%#v\n", c)
}
