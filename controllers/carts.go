package controllers

import (
	"net/http"

	"github.com/ardin2001/go_mini-capstone/helpers"
	"github.com/ardin2001/go_mini-capstone/middlewares"
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/services"
	"github.com/labstack/echo/v4"
)

type CartInterfaceC interface {
	GetCartsController(c echo.Context) error
	CreateCartController(c echo.Context) error
	UpdateCartController(c echo.Context) error
	DeleteCartController(c echo.Context) error
}

type CartStructC struct {
	cartS services.CartInterfaceS
}

func NewCartControllers(cartS services.CartInterfaceS) CartInterfaceC {
	return &CartStructC{
		cartS,
	}
}

func (cc *CartStructC) GetCartsController(c echo.Context) error {
	data, err := middlewares.AdminVerification(c)
	if !err {
		return data
	}
	carts, check := cc.cartS.GetCartsService()
	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    carts,
		Message: "Successfull get carts account",
		Status:  true,
	})
}

func (cc *CartStructC) CreateCartController(c echo.Context) error {
	cart := models.Cart{}
	c.Bind(&cart)

	_, check := cc.cartS.CreateCartService(&cart)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    cart,
		Message: "Successfull create cart account",
		Status:  true,
	})
}

func (cc *CartStructC) UpdateCartController(c echo.Context) error {
	data, err := middlewares.AdminVerification(c)
	if !err {
		return data
	}

	id := c.Param("id")
	cart := models.Cart{}
	c.Bind(&cart)

	dataCart, check := cc.cartS.UpdateCartService(&cart, id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    dataCart,
		Message: "Successfull update cart account",
		Status:  true,
	})
}

func (cc *CartStructC) DeleteCartController(c echo.Context) error {
	id := c.Param("id")
	check := cc.cartS.DeleteCartService(id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    id,
		Message: "Successfull delete cart account",
		Status:  true,
	})
}
