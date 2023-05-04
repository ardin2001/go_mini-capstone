package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string        `json:"nama" form:"nama"`
	Email        string        `json:"email" form:"email"`
	Password     string        `json:"password" form:"password"`
	Role         string        `json:"role" form:"role"`
	No_HP        string        `json:"no_hp" form:"no_hp"`
	Carts        []Cart        `gorm:"foreignKey:UserId"`
	Transactions []Transaction `gorm:"foreignKey:UserId"`
}

type JwtCustomClaims struct {
	ID   uint   `json:"userId" form:"userId"`
	Nama string `json:"nama" form:"nama"`
	Role string `json:"role" form:"role"`
	jwt.RegisteredClaims
}

type UserResponse struct {
	ID    uint   `json:"id" form:"id"`
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"email"`
	Role  string `json:"role" form:"role"`
	Token string `json:"token" form:"token"`
}
