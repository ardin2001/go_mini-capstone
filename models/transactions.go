package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId             uint                `json:"user_id" form:"user_id"`
	Status             bool                `json:"status" form:"status"`
	BuktiTransaksi     string              `json:"bukti_tf" form:"bukti_tf"`
	User               User                `json:"user"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionId"`
}

type TransactionResult struct {
	UserId         uint   `json:"user_id" form:"user_id"`
	Status         bool   `json:"status" form:"status"`
	BuktiTransaksi string `json:"bukti_tf" form:"bukti_tf"`
	JumlahBarang   string `json:"jumlah_barang" form:"jumlah_barang"`
	TotalHarga     string `json:"total_harga" form:"total_harga"`
}
