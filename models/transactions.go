package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId             uint                `json:"user_id" form:"user_id"`
	Status             bool                `json:"status" form:"status"`
	Alamat             string              `json:"alamat" form:"alamat"`
	Ongkir             int                 `json:"ongkir" form:"ongkir"`
	Ekspedisi          string              `json:"ekspedisi" form:"ekspedisi"`
	BuktiTransaksi     string              `json:"bukti_tf" form:"bukti_transaksi"`
	JumlahBarang       int                 `json:"jumlah_barang" form:"jumlah_barang" gorm:"-:all"`
	TotalHarga         int                 `json:"total_harga" form:"total_harga" gorm:"-:all"`
	Tujuan             string              `json:"tujuan" form:"tujuan" gorm:"-:all"`
	User               User                `json:"user"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionId"`
}
