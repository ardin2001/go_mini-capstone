package models

import (
	"gorm.io/gorm"
)

type TransactionDetail struct {
	gorm.Model
	ProductId     uint        `json:"product_id" form:"product_id"`
	TransactionId uint        `json:"transaction_id" form:"transaction_id"`
	Jumlah        int         `json:"jumlah" form:"jumlah"`
	CartId        int         `json:"cart_id" gorm:"-"`
	Product       Product     `json:"product"`
	Transaction   Transaction `json:"transaction"`
}
