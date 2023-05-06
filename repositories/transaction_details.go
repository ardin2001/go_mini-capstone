package repositories

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"gorm.io/gorm"
)

type TransactionDetailInterfaceR interface {
	GetTransactionDetailsRepository(id string) ([]models.TransactionDetail, error)
	CreateTransactionDetailRepository(transactionDetail *[]models.TransactionDetail) (*[]models.TransactionDetail, error)
}

type TransactionDetailStructR struct {
	DB *gorm.DB
}

func NewTransactionDetailRepositories(db *gorm.DB) TransactionDetailInterfaceR {
	return &TransactionDetailStructR{
		DB: db,
	}
}

func (tdr *TransactionDetailStructR) GetTransactionDetailsRepository(id string) ([]models.TransactionDetail, error) {
	var transaction_details []models.TransactionDetail
	var check error
	if id == "" {
		check = tdr.DB.Preload("Product").Find(&transaction_details).Error
	} else {
		check = tdr.DB.Where("user_id", id).Preload("Product").Find(&transaction_details).Error
	}
	if check != nil {
		return nil, check
	}

	return transaction_details, check
}

func (cr *TransactionDetailStructR) CreateTransactionDetailRepository(transactionDetail *[]models.TransactionDetail) (*[]models.TransactionDetail, error) {
	check := cr.DB.Create(transactionDetail).Error
	if check != nil {
		return nil, check
	}
	return transactionDetail, check
}
