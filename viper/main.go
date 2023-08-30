package main

import (
	_"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"app_name"`
	LogLevel string `mapstructure:"log_level"`
	MysqlInfo MysqlConfig `mapstructure:"mysql"`
	RedisInfo RedisConfig `mapstructure:"redis"`
}

type MysqlConfig struct {
	Ip string `mapstructure:"ip"`
	Port int `mapstructure:"port"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Ip string `mapstructure:"ip"`
	Port int `mapstructure:"port"`
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config failed, err: ", err)
		return
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("unmarshal failed, err: ", err)
		return
	}
	fmt.Printf("config: %v\n", config)
}