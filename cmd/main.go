package main

import (
	"service/config"
	"service/dao"
	"service/log"
	"service/routers"
)

func init() {
	config.Init()
	log.Init(log.LogMode(config.LogConfigMode))
	dao.Init()
	routers.Init()

}

func main() {

	return
}
