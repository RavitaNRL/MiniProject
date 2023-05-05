package models

import "time"

type status string

const (
	PENDING string = "PENDING"
	PROCESS string = "PROCESS"
	FINISH  string = "FINISH"
)

type Order struct {
	ID             int       `json:"id" gorm:"primary_key" form:"id"`
	ID_user        int       `json:"id_user" form:"id_user" gorm:"autoIncrement"`
	ID_baju        int       `json:"id_baju" form:"id_baju" gorm:"autoIncrement"`
	Jumlah         int       `json:"jumlah" form:"jumlah"`
	Tanggal_order  string    `json:"tanggal_order" form:"tanggal_order"`
	Total_harga    int       `json:"total_harga" form:"total_harga"`
	Status         status    `json:"status" form:"status" gorm:"type:enum('PENDING', 'PROCESS', 'FINISH')";"column:status"`
	Estimasi_Waktu time.Time `json:"estimasi_waktu" form:"estimasi_waktu"`
}
