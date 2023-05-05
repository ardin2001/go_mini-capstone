package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId             uint                `json:"user_id" form:"user_id"`
	Status             bool                `json:"status" form:"status"`
	BuktiTransaksi     string              `json:"bukti_tf" form:"bukti_tf"`
	JumlahBarang       int                 `json:"jumlah_barang" form:"jumlah_barang" gorm:"-:all"`
	TotalHarga         int                 `json:"total_harga" form:"total_harga" gorm:"-:all"`
	User               User                `json:"user"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionId"`
}
