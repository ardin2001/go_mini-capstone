package repositories

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"gorm.io/gorm"
)

type TransactionInterfaceR interface {
	GetTransactionsRepository(id string) ([]models.Transaction, error)
	// GetTransactionRepository(id, user_id string) (*models.Transaction, error)
	CreateTransactionRepository(Transaction *models.Transaction) (*models.Transaction, error)
	// DeleteTransactionRepository(id, id_user string) error
	// UpdateTransactionRepository(TransactionId *models.Transaction, id string) (*models.Transaction, error)
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
		check = tr.DB.Where("user_id", id).Preload("User").Preload("TransactionDetails").Find(&transactions).Error
	}
	if check != nil {
		return nil, check
	}

	return transactions, check
}

// func (cr *TransactionStructR) GetTransactionRepository(id, user_id string) (*models.Transaction, error) {
// 	var cart models.Transaction
// 	check := cr.DB.Where("user_id", user_id).Preload("User").Preload("Product").First(&cart, id).Error
// 	if check != nil {
// 		return nil, check
// 	}
// 	return &cart, check
// }

func (cr *TransactionStructR) CreateTransactionRepository(transaction *models.Transaction) (*models.Transaction, error) {
	check := cr.DB.Save(transaction).Error
	if check != nil {
		return nil, check
	}
	return transaction, check
}

// func (cr *TransactionStructR) DeleteCartRepository(id, user_id string) error {
// 	if err := cr.DB.Where("id = ?", id).Take(&models.Transaction{}).Error; err != nil {
// 		return errors.New("not_found")
// 	}

// 	check := cr.DB.Where("user_id", user_id).Delete(&models.Transaction{}, &id).Error
// 	fmt.Println(check, user_id)
// 	if check != nil {
// 		return errors.New("protected")
// 	}
// 	return check
// }

// func (cr *TransactionStructR) UpdateCartRepository(cartId *models.Transaction, id string) (*models.Transaction, error) {
// 	check := cr.DB.Save(cartId).Error
// 	if check != nil {
// 		return nil, check
// 	}
// 	return cartId, check
// }
