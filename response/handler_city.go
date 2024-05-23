package response

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/logic"
)

func GetCityList(c *gin.Context) {
	cityList := logic.CityList()
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"citylist": cityList,
		},
	})
}

// 公钥
type AddressRequest struct {
	Address string `json:"address"`
}

func ValidateCity(c *gin.Context) {

	var request AddressRequest
	// 将请求体绑定到结构体
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address := request.Address

	// 验证地址是否为有效的以太坊地址
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Ethereum address"})
		return
	}

	// 将字符串形式的地址转换为 common.Address 类型
	ethAddress := common.HexToAddress(address)

	ok := logic.ValidateHasAllCity(ethAddress)

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"allow": ToString(ok),
		},
	})
}

func ToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
