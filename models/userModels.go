package models

type User struct {
	ID        int    `json:"id" gorm:"primary_key" form:"id"`
	Nama      string `json:"nama" form:"nama"`
	No_Telp   string `json:"no_telp" form:"no_telp"`
	Alamat    string `json:"alamat" form:"alamat"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	User_type string `json:"user_type" form:"user_type"`
	Token     string `json:"token" form:"token"`
}
