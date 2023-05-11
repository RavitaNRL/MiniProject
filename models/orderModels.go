package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	User           User   `json:"user"  gorm:"foreignKey:ID_user;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ID_user        int    `json:"id_user" form:"id_user"`
	Baju           Baju   `json:"baju" gorm:"foreignKey:ID_baju;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ID_baju        int    `json:"id_baju" form:"id_baju"`
	Jumlah         int    `json:"jumlah" form:"jumlah"`
	Total_harga    int    `json:"total_harga" form:"total_harga"`
	Status         string `json:"status_pesanan" form:"status_pesanan"`
	Tanggal_order  string `json:"tanggal_order" form:"tanggal_order"`
	Estimasi_Waktu string `json:"estimasi_waktu" form:"estimasi_waktu"`
}
