package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId             uint                `json:"user_id" form:"user_id"`
	Status             bool                `json:"status" form:"status"`
	BuktiTransaksi     string              `json:"bukti_tf" form:"bukti_tf"`
	TotalHarga         int                 `json:"total_harga" gorm:"-"`
	User               User                `json:"user"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionId"`
}
