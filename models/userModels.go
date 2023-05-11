package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Alamat    string `json:"alamat" form:"alamat"`
	No_Hp     string `json:"no_hp" form:"no_hp"`
	User_type string `json:"user_type" form:"user_type"`
	Password  string `json:"password" form:"password"`
	Token     string `json:"token" form:"token"`
}
