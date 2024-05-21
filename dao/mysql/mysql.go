package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"service/config"
	"service/models"
)

var db *gorm.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlPort,
		config.MysqlDbname,
	)

	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//err = db.Table(models.City{}.TableName()).AutoMigrate(&models.City{})
	//if err != nil {
	//	return err
	//}
	_ = db.Migrator().CreateTable(models.City{})
	return
}
