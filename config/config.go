package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"service/log"
)

var (
	ServerName    string
	ServerPort    string
	LogConfigMode int
	MysqlHost     string
	MysqlPort     int
	MysqlUser     string
	MysqlPassword string
	MysqlDbname   string
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

func LoadConfig() {
	ServerName = viper.GetString("server.name")
	ServerPort = viper.GetString("server.port")
	LogConfigMode = viper.GetInt("logConfig.mode")
	MysqlHost = viper.GetString("mysql.host")
	MysqlPort = viper.GetInt("mysql.port")
	MysqlUser = viper.GetString("mysql.user")
	MysqlPassword = viper.GetString("mysql.password")
	MysqlDbname = viper.GetString("mysql.dbname")
}
