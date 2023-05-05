package models

type Baju struct {
	ID         int    `json:"id" gorm:"primary_key" form:"id"`
	Nama_baju  string `json:"nama_baju" form:"nama_baju"`
	Jenis_baju string `json:"jenis_baju" form:"jenis_baju"`
	Harga_baju int    `json:"harga_baju" form:"harga_baju"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
}
