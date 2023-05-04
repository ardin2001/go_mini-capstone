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
	// GetTransactionDetailController(c echo.Context) error
	CreateTransactionDetailController(c echo.Context) error
	// UpdateTransactionDetailController(c echo.Context) error
	// DeleteTransactionDetailController(c echo.Context) error
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

// func (cc *CartStructC) DeleteCartController(c echo.Context) error {
// 	id := c.Param("id")
// 	user := middlewares.GetDataJWT(c)
// 	user_id := strconv.Itoa(int(user.ID))
// 	check := cc.cartS.DeleteCartService(id, user_id)

// 	if check != nil {
// 		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
// 			Data:    nil,
// 			Message: "err()",
// 			Status:  false,
// 		})
// 	}
// 	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
// 		Data:    id,
// 		Message: "Successfull delete cart account",
// 		Status:  true,
// 	})
// }
