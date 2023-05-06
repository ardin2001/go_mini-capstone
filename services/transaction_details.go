package services

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"
)

type TransactionDetailInterfaceS interface {
	GetTransactionDetailsService(id string) ([]models.TransactionDetail, error)
	CreateTransactionDetailService(transactionDetail *[]models.TransactionDetail) (*[]models.TransactionDetail, error)
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

func (cs *TransactionDetailStructS) CreateTransactionDetailService(transactionDetail *[]models.TransactionDetail) (*[]models.TransactionDetail, error) {
	transactionDetailR, err := cs.transactionDetailR.CreateTransactionDetailRepository(transactionDetail)
	if err != nil {
		return nil, err
	}

	return transactionDetailR, nil
}
