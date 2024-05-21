package logic

import "service/models"

func FindCityAll() {
	a := models.GetAllCityNames(models.City{}.Model)

}
