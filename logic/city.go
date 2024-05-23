package logic

import (
	"github.com/ethereum/go-ethereum/common"
	"service/eth"
)

func CityList() []string {
	cityExist := []string{"ShangHai", "BeiJing", "ChengDu", "WuHan", "ShenZhen"}
	return cityExist
}

func ValidateHasAllCity(address common.Address) bool {

	cityExist := []string{"ShangHai", "BeiJing", "ChengDu", "WuHan", "ShenZhen"}
	cityList := eth.ContractQueryCityMint(address)

	cityMap := make(map[string]bool)
	for _, city := range cityList {
		cityMap[city] = true
	}

	// Check if each city in cityExist is present in cityMap
	for _, city := range cityExist {
		if !cityMap[city] {
			return false
		}
	}
	return true
}
