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

type CartInterfaceC interface {
	GetCartsController(c echo.Context) error
	GetCartController(c echo.Context) error
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
	var carts []models.Cart
	var check error
	_, err := middlewares.AdminVerification(c)
	if err {
		carts, check = cc.cartS.GetCartsService("")
	} else {
		data := middlewares.GetDataJWT(c)
		id := strconv.Itoa(int(data.ID))
		carts, check = cc.cartS.GetCartsService(id)
	}

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}

	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    carts,
		Message: "Successfull get carts account",
		Status:  true,
	})
}

func (cc *CartStructC) GetCartController(c echo.Context) error {
	data := middlewares.GetDataJWT(c)
	user_id := strconv.Itoa(int(data.ID))
	id := c.Param("id")
	carts, check := cc.cartS.GetCartService(id, user_id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
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
	data := middlewares.GetDataJWT(c)
	cart.UserId = data.ID
	_, check := cc.cartS.CreateCartService(&cart)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
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
	id := c.Param("id")
	cart := models.Cart{}
	c.Bind(&cart)

	user := middlewares.GetDataJWT(c)
	user_id := strconv.Itoa(int(user.ID))
	dataCart, check := cc.cartS.UpdateCartService(&cart, id, user_id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
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
	user := middlewares.GetDataJWT(c)
	user_id := strconv.Itoa(int(user.ID))
	check := cc.cartS.DeleteCartService(id, user_id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    id,
		Message: "Successfull delete cart account",
		Status:  true,
	})
}
