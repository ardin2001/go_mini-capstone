package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID        uint    `json:"cart_id" gorm:"primaryKey"`
	ProductId uint    `json:"product_id" form:"product_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	Jumlah    int     `json:"jumlah" form:"jumlah"`
	User      User    `json:"user"`
	Product   Product `json:"product"`
}
