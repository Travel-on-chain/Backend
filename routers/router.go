package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	// 支持跨域
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true, // 允许携带 Cookie
	}))

	v1 := r.Group("/api")

	// 检测用户是否登录
	v1.GET("/check-login", nil)
	// 登录
	v1.GET("/nonce", nil)
	v1.POST("/login", nil)
	// 获取历史位置
	v1.POST("/location", nil)

	return nil
}
