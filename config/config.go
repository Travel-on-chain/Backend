package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"service/log"
)

func Init() {
	// 读取 配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	// 监控 配置文件更改
	viper.WatchConfig()
	viper.OnConfigChange(
		func(in fsnotify.Event) {
			log.Error("The configuration file has been modified.")
		})
	return
}
