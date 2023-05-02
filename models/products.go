package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Nama      string `json:"nama" form:"nama"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Berat     int    `json:"berat" form:"berat"`
	Harga     int    `json:"harga" form:"harga"`
	Status    string `json:"status" form:"status"`
	Gambar    string `json:"gambar" form:"gambar"`
	Carts     []Cart `gorm:"foreignKey:ProductId"`
}
