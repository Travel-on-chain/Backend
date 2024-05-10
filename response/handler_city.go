package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCityList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"citylist": "北京,上海,成都,西安,深圳,武汉,哈尔滨",
		},
	})
}
