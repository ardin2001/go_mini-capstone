package services

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"
)

type TransactionInterfaceS interface {
	GetTransactionsService(id string) ([]models.Transaction, error)
	GetTransactionService(id, user_id string) (*models.Transaction, error)
	CreateTransactionService(Transaction *models.Transaction) (*models.Transaction, error)
	UpdateTransactionService(TransactionId *models.Transaction, id, user_id, user_role string) (*models.Transaction, error)
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

	for i := range transactions {
		for j := range transactions[i].TransactionDetails {
			transactions[i].JumlahBarang += transactions[i].TransactionDetails[j].Jumlah
			transactions[i].TotalHarga += transactions[i].TransactionDetails[j].Product.Harga * transactions[i].TransactionDetails[j].Jumlah
		}
	}

	return transactions, nil
}

func (ts *TransactionStructS) GetTransactionService(id, user_id string) (*models.Transaction, error) {
	transaction, err := ts.transactionR.GetTransactionRepository(id, user_id)
	if err != nil {
		return nil, err
	}

	for j := range transaction.TransactionDetails {
		transaction.JumlahBarang += transaction.TransactionDetails[j].Jumlah
		transaction.TotalHarga += transaction.TransactionDetails[j].Product.Harga * transaction.TransactionDetails[j].Jumlah
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

func (ts *TransactionStructS) UpdateTransactionService(transactionId *models.Transaction, id, user_id, user_role string) (*models.Transaction, error) {
	getTransactionId, err := ts.transactionR.GetTransactionRepository(id, user_id)

	if err != nil {
		return nil, err
	}

	if user_role == "admin" {
		getTransactionId.Status = transactionId.Status
	} else {
		if transactionId.BuktiTransaksi != "" {
			getTransactionId.BuktiTransaksi = transactionId.BuktiTransaksi
		}
		if transactionId.Alamat != "" {
			getTransactionId.Alamat = transactionId.Alamat
		}
		if transactionId.Ongkir != 0 {
			getTransactionId.Ongkir = transactionId.Ongkir
		}
		if transactionId.Ekspedisi != "" {
			getTransactionId.Ekspedisi = transactionId.Ekspedisi
		}
	}

	cart, err := ts.transactionR.UpdateTransactionRepository(getTransactionId, id)
	if err != nil {
		return nil, err
	}

	for j := range getTransactionId.TransactionDetails {
		getTransactionId.JumlahBarang += getTransactionId.TransactionDetails[j].Jumlah
		getTransactionId.TotalHarga += getTransactionId.TransactionDetails[j].Product.Harga * getTransactionId.TransactionDetails[j].Jumlah
	}

	return cart, nil
}
