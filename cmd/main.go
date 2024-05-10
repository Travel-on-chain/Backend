package cmd

import (
	"service/config"
	"service/log"
	"service/routers"
)

func init() {
	config.Init()
	log.Init(log.LogMode(config.LogConfigMode))
	routers.Init()
}

func main() {

	return
}
