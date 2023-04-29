package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	No_HP    string `json:"no_hp" form:"no_hp"`
	Alamat   string `json:"alamat" form:"alamat"`
}

type JwtCustomClaims struct {
	ID   uint   `json:"userId" form:"userId"`
	Name string `json:"name" form:"name"`
	Role string `json:"role" form:"role"`
	jwt.RegisteredClaims
}

type UserResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Role  string `json:"role" form:"role"`
	Token string `json:"token" form:"token"`
}
