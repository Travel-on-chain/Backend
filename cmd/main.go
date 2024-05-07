package cmd

import (
	"github.com/spf13/viper"
	"service/config"
	"service/log"
	"service/routers"
)

func init() {
	config.Init()
	log.Init(log.LogMode(viper.GetInt("LogConfig.mode")))
	routers.Init()
}

func main() {

	return
}
