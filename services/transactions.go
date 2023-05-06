package services

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"
)

type TransactionInterfaceS interface {
	GetTransactionsService(id string) ([]models.Transaction, error)
	GetTransactionService(id, user_id string) (*models.Transaction, error)
	CreateTransactionService(Transaction *models.Transaction) (*models.Transaction, error)
	// UpdateTransactionService(TransactionId *models.Transaction, id, user_id string) (*models.Transaction, error)
}

type TransactionStructS struct {
	transactionR repositories.TransactionInterfaceR
}

func NewTransactionServices(transactionR repositories.TransactionInterfaceR) TransactionInterfaceS {
	return &TransactionStructS{
		transactionR: transactionR,
	}
}

func (ts *TransactionStructS) GetTransactionsService(id string) ([]models.Transaction, error) {
	transactions, err := ts.transactionR.GetTransactionsRepository(id)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (ts *TransactionStructS) GetTransactionService(id, user_id string) (*models.Transaction, error) {
	transaction, err := ts.transactionR.GetTransactionRepository(id, user_id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (cs *TransactionStructS) CreateTransactionService(cart *models.Transaction) (*models.Transaction, error) {
	transactionR, err := cs.transactionR.CreateTransactionRepository(cart)
	if err != nil {
		return nil, err
	}

	return transactionR, nil
}

// func (cs *CartStructS) UpdateCartService(cartId *models.Transaction, id, user_id string) (*models.Transaction, error) {
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
