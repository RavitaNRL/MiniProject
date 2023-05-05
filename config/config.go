package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"miniProject/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(localhost:3306)/mini-project?charset=utf8mb4&parseTime=True&loc=Local"
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
