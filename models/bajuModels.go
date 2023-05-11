package models

import "gorm.io/gorm"

type Baju struct {
	gorm.Model
	Nama_Pembeli string `json:"nama_pembeli" form:"nama_pembeli"`
	Jenis_Baju   string `json:"jenis_baju" form:"jenis_baju"`
	Deksripsi    string `json:"deksripsi_baju" form:"deksripsi_baju"`
	Ketarangan   string `json:"keterangan" form:"keterangan"`
	// Orders       []Order `json:"orders" gorm:"foreignKey:ID_baju;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
