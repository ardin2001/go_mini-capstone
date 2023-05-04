package services

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"
)

type TransactionDetailInterfaceS interface {
	GetTransactionDetailsService(id string) ([]models.TransactionDetail, error)
	// GetTransactionDetailService(id, user_id string) (*models.TransactionDetail, error)
	CreateTransactionDetailService(transactionDetail *[]models.TransactionDetail) (*[]models.TransactionDetail, error)
	// UpdateTransactionDetailService(TransactionDetailId *models.TransactionDetail, id, user_id string) (*models.TransactionDetail, error)
	// DeleteTransactionDetailService(id, user_id string) error
}

type TransactionDetailStructS struct {
	transactionDetailR repositories.TransactionDetailInterfaceR
}

func NewTransactionDetailServices(transactionDetailR repositories.TransactionDetailInterfaceR) TransactionDetailInterfaceS {
	return &TransactionDetailStructS{
		transactionDetailR: transactionDetailR,
	}
}

func (tds *TransactionDetailStructS) GetTransactionDetailsService(id string) ([]models.TransactionDetail, error) {
	transactionDetails, err := tds.transactionDetailR.GetTransactionDetailsRepository(id)
	if err != nil {
		return nil, err
	}

	return transactionDetails, nil
}

// func (cs *TransactionDetailStructS) GetCartService(id, user_id string) (*models.TransactionDetail, error) {
// 	cart, err := cs.cartR.GetCartRepository(id, user_id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cart, nil
// }

func (cs *TransactionDetailStructS) CreateTransactionDetailService(transactionDetail *[]models.TransactionDetail) (*[]models.TransactionDetail, error) {
	transactionDetailR, err := cs.transactionDetailR.CreateTransactionDetailRepository(transactionDetail)
	if err != nil {
		return nil, err
	}

	return transactionDetailR, nil
}

// func (cs *TransactionDetailStructS) UpdateCartService(cartId *models.TransactionDetail, id, user_id string) (*models.TransactionDetail, error) {
// 	getCartId, err := cs.cartR.GetCartRepository(id, user_id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if cartId.Jumlah != 0 {
// 		getCartId.Jumlah = cartId.Jumlah
// 	}

// 	cart, err := cs.cartR.UpdateCartRepository(getCartId, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cart, nil
// }

// func (cs *TransactionDetailStructS) DeleteCartService(id, user_id string) error {
// 	err := cs.cartR.DeleteCartRepository(id, user_id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
