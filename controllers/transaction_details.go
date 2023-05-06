package controllers

import (
	"net/http"
	"strconv"

	"github.com/ardin2001/go_mini-capstone/helpers"
	"github.com/ardin2001/go_mini-capstone/middlewares"
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/services"
	"github.com/labstack/echo/v4"
)

type TransactionDetailInterfaceC interface {
	GetTransactionDetailsController(c echo.Context) error
	CreateTransactionDetailController(c echo.Context) error
}

type TransactionDetailStructC struct {
	transactionDetailS services.TransactionDetailInterfaceS
}

func NewTransactionDetailControllers(transactionDetailS services.TransactionDetailInterfaceS) TransactionDetailInterfaceC {
	return &TransactionDetailStructC{
		transactionDetailS,
	}
}

func (tdc *TransactionDetailStructC) GetTransactionDetailsController(c echo.Context) error {
	var transactionDetails []models.TransactionDetail
	var check error
	_, err := middlewares.AdminVerification(c)
	if err {
		transactionDetails, check = tdc.transactionDetailS.GetTransactionDetailsService("")
	} else {
		data := middlewares.GetDataJWT(c)
		id := strconv.Itoa(int(data.ID))
		transactionDetails, check = tdc.transactionDetailS.GetTransactionDetailsService(id)
	}

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    transactionDetails,
		Message: "Successfull get transaction",
		Status:  true,
	})
}

func (cc *TransactionDetailStructC) CreateTransactionDetailController(c echo.Context) error {
	transaction_details := []models.TransactionDetail{}
	c.Bind(&transaction_details)
	_, check := cc.transactionDetailS.CreateTransactionDetailService(&transaction_details)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    transaction_details,
		Message: "Successfull create cart account",
		Status:  true,
	})
}
