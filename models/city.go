package models

import "gorm.io/gorm"

// City 代表城市模型及其属性
type City struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null" json:"name"`
}

// TableName 设置该结构类型的插入表名
func (City) TableName() string {
	return "City"
}

// GetAllCityNames 获取数据库中所有城市名称
// @Summary 获取所有城市名称
// @Description 获取数据库中所有城市名称的列表
// @Tags 城市
// @Accept  json
// @Produce  json
// @Param db query string true "数据库连接"
// @Success 200 {array} string "城市名称列表"
// @Failure 500 {object} string "内部服务器错误"
// @Router /cities [get]
func GetAllCityNames(db *gorm.DB) ([]string, error) {
	var cities []City
	var cityNames []string
	if err := db.Select("name").Find(&cities).Error; err != nil {
		return nil, err
	}
	for _, city := range cities {
		cityNames = append(cityNames, city.Name)
	}
	return cityNames, nil
}

// IsCityExists 检查特定城市是否存在于数据库中
// @Summary 检查城市是否存在
// @Description 验证数据库中是否存在具有给定名称的城市
// @Tags 城市
// @Accept  json
// @Produce  json
// @Param db query string true "数据库连接"
// @Param cityName path string true "城市名称"
// @Success 200 {boolean} bool "城市存在"
// @Failure 404 {object} string "城市不存在"
// @Failure 500 {object} string "内部服务器错误"
// @Router /cities/{cityName} [get]
func IsCityExists(db *gorm.DB, cityName string) (bool, error) {
	var city City
	if err := db.Where("name = ?", cityName).First(&city).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
