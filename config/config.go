package config

import (
	"Project-Mini/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(localhost:3306)/project_mini?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigration()
}

func initMigration() {
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Baju{})
	DB.AutoMigrate(&models.User{})
}
