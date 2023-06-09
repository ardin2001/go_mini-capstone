package utils

import (
	"fmt"

	"github.com/ardin2001/go_mini-capstone/configs"
	"github.com/ardin2001/go_mini-capstone/models"
)

func UserMigrate() {
	DB, err := configs.InitDB()
	if err != nil {
		fmt.Println("Failed connect to database : ", err.Error())
		return
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Transaction{})
	DB.AutoMigrate(&models.TransactionDetail{})
}
