package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type JwtCustomClaims struct {
	ID   uint   `json:"userId" form:"userId"`
	Name string `json:"name" form:"name"`
	jwt.RegisteredClaims
}

type UserResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
