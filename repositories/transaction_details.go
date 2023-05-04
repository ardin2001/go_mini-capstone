package repositories

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"gorm.io/gorm"
)

type TransactionDetailInterfaceR interface {
	GetTransactionDetailsRepository(id string) ([]models.TransactionDetail, error)
	// GetTransactionDetailRepository(id, user_id string) (*models.TransactionDetail, error)
	CreateTransactionDetailRepository(transactionDetail *[]models.TransactionDetail) (*[]models.TransactionDetail, error)
	// DeleteTransactionDetailRepository(id, id_user string) error
	// UpdateTransactionDetailRepository(TransactionDetailId *models.TransactionDetail, id string) (*models.TransactionDetail, error)
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
		check = tdr.DB.Preload("Product").Preload("Transaction").Find(&transaction_details).Error
	} else {
		check = tdr.DB.Where("user_id", id).Preload("Product").Preload("Transaction").Find(&transaction_details).Error
	}
	if check != nil {
		return nil, check
	}

	return transaction_details, check
}

// func (cr *CartStructR) GetCartRepository(id, user_id string) (*models.TransactionDetail, error) {
// 	var cart models.TransactionDetail
// 	check := cr.DB.Where("user_id", user_id).Preload("User").Preload("Product").First(&cart, id).Error
// 	if check != nil {
// 		return nil, check
// 	}
// 	return &cart, check
// }

func (cr *TransactionDetailStructR) CreateTransactionDetailRepository(transactionDetail *[]models.TransactionDetail) (*[]models.TransactionDetail, error) {
	check := cr.DB.Create(transactionDetail).Error
	if check != nil {
		return nil, check
	}
	return transactionDetail, check
}

// func (cr *CartStructR) DeleteCartRepository(id, user_id string) error {
// 	if err := cr.DB.Where("id = ?", id).Take(&models.TransactionDetail{}).Error; err != nil {
// 		return errors.New("not_found")
// 	}

// 	check := cr.DB.Where("user_id", user_id).Delete(&models.TransactionDetail{}, &id).Error
// 	fmt.Println(check, user_id)
// 	if check != nil {
// 		return errors.New("protected")
// 	}
// 	return check
// }

// func (cr *CartStructR) UpdateCartRepository(cartId *models.TransactionDetail, id string) (*models.TransactionDetail, error) {
// 	check := cr.DB.Save(cartId).Error
// 	if check != nil {
// 		return nil, check
// 	}
// 	return cartId, check
// }
