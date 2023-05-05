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

type TransactionInterfaceC interface {
	GetTransactionsController(c echo.Context) error
	// GetTransactionController(c echo.Context) error
	CreateTransactionController(c echo.Context) error
	// UpdateTransactionController(c echo.Context) error
}

type TransactionStructC struct {
	transactionS       services.TransactionInterfaceS
	transactionDetailS services.TransactionDetailInterfaceS
}

func NewTransactionControllers(transactionS services.TransactionInterfaceS, transactionDetailS services.TransactionDetailInterfaceS) TransactionInterfaceC {
	return &TransactionStructC{
		transactionS,
		transactionDetailS,
	}
}

func (tc *TransactionStructC) GetTransactionsController(c echo.Context) error {
	var transactions []models.Transaction
	var check error
	_, err := middlewares.AdminVerification(c)
	if err {
		transactions, check = tc.transactionS.GetTransactionsService("")
	} else {
		data := middlewares.GetDataJWT(c)
		id := strconv.Itoa(int(data.ID))
		transactions, check = tc.transactionS.GetTransactionsService(id)
	}

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    transactions,
		Message: "Successfull get transaction",
		Status:  true,
	})
}

// func (cc *CartStructC) GetCartController(c echo.Context) error {
// 	data := middlewares.GetDataJWT(c)
// 	user_id := strconv.Itoa(int(data.ID))
// 	id := c.Param("id")
// 	carts, check := cc.cartS.GetCartService(id, user_id)

// 	if check != nil {
// 		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
// 			Data:    nil,
// 			Message: "err()",
// 			Status:  false,
// 		})
// 	}

// 	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
// 		Data:    carts,
// 		Message: "Successfull get carts account",
// 		Status:  true,
// 	})
// }

func (tc *TransactionStructC) CreateTransactionController(c echo.Context) error {
	transactions := models.Transaction{}
	transaction_details := []models.TransactionDetail{}

	c.Bind(&transaction_details)
	data := middlewares.GetDataJWT(c)
	transactions.UserId = data.ID
	_, check := tc.transactionS.CreateTransactionService(&transactions)
	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}

	for i := range transaction_details {
		transaction_details[i].TransactionId = transactions.ID
	}

	_, check2 := tc.transactionDetailS.CreateTransactionDetailService(&transaction_details)
	if check2 != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check2.Error(),
			Status:  false,
		})
	}

	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    transactions,
		Message: "Successfull create transaction_details",
		Status:  true,
	})
}

// func (cc *CartStructC) UpdateCartController(c echo.Context) error {
// 	id := c.Param("id")
// 	cart := models.Cart{}
// 	c.Bind(&cart)

// 	user := middlewares.GetDataJWT(c)
// 	user_id := strconv.Itoa(int(user.ID))
// 	dataCart, check := cc.cartS.UpdateCartService(&cart, id, user_id)

// 	if check != nil {
// 		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
// 			Data:    nil,
// 			Message: "err()",
// 			Status:  false,
// 		})
// 	}
// 	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
// 		Data:    dataCart,
// 		Message: "Successfull update cart account",
// 		Status:  true,
// 	})
// }
