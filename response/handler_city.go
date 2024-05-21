package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCityList(c *gin.Context) {
	cityList := logic.
		c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"citylist": cityList,
		},
	})
}
