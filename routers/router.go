package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"service/config"
	"service/response"
)

func Init() *gin.Engine {
	// 支持跨域
	service := gin.Default()
	service.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true, // 允许携带 Cookie
	}))

	v1 := service.Group("/user")

	// 登录相关
	//v1.GET("/login", nil)
	//v1.GET("/signin", nil)
	//v1.GET("/logout", nil)
	//// 获取历史位置
	//v1.POST("/location", nil)
	//v1.GET("/lots", nil)
	//
	v1.POST("/citylist", response.GetCityList)

	err := service.Run(config.ServerPort)
	if err != nil {
		return nil
	}
	return nil
}
