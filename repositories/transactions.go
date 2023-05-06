package repositories

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"gorm.io/gorm"
)

type TransactionInterfaceR interface {
	GetTransactionsRepository(id string) ([]models.Transaction, error)
	GetTransactionRepository(id, user_id string) (*models.Transaction, error)
	CreateTransactionRepository(Transaction *models.Transaction) (*models.Transaction, error)
	UpdateTransactionRepository(TransactionId *models.Transaction, id string) (*models.Transaction, error)
}

type TransactionStructR struct {
	DB *gorm.DB
}

func NewTransactionRepositories(db *gorm.DB) TransactionInterfaceR {
	return &TransactionStructR{
		DB: db,
	}
}

func (tr *TransactionStructR) GetTransactionsRepository(id string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	var check error
	if id == "" {
		check = tr.DB.Preload("User").Preload("TransactionDetail").Find(&transactions).Error
	} else {
		check = tr.DB.Where("user_id", id).Preload("User").Preload("TransactionDetails.Product").Find(&transactions).Error
	}
	if check != nil {
		return nil, check
	}

	return transactions, check
}

func (cr *TransactionStructR) GetTransactionRepository(id, user_id string) (*models.Transaction, error) {
	var transaction models.Transaction
	var check error
	if user_id == "1" {
		check = cr.DB.Preload("User").Preload("TransactionDetails.Product").First(&transaction, id).Error
	} else {
		check = cr.DB.Where("user_id", user_id).Preload("User").Preload("TransactionDetails.Product").First(&transaction, id).Error
	}

	if check != nil {
		return nil, check
	}
	return &transaction, check
}

func (cr *TransactionStructR) CreateTransactionRepository(transaction *models.Transaction) (*models.Transaction, error) {
	check := cr.DB.Save(transaction).Error
	if check != nil {
		return nil, check
	}
	return transaction, check
}

func (cr *TransactionStructR) UpdateTransactionRepository(transactionId *models.Transaction, id string) (*models.Transaction, error) {
	check := cr.DB.Save(transactionId).Error
	if check != nil {
		return nil, check
	}
	return transactionId, check
}
