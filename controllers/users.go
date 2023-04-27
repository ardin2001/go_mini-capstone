package controllers

import (
	"net/http"

	"github.com/ardin2001/go_mini-capstone/helpers"
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/services"
	"github.com/labstack/echo/v4"
)

type UserInterfaceC interface {
	LoginUserController(c echo.Context) error
	GetUsersController(c echo.Context) error
	GetUserController(c echo.Context) error
	CreateUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
	DeleteUserController(c echo.Context) error
}

type UserStructC struct {
	userS services.UserInterfaceS
}

func NewUserControllers(userS services.UserInterfaceS) UserInterfaceC {
	return &UserStructC{
		userS,
	}
}

func (u *UserStructC) LoginUserController(c echo.Context) error {
	err := true

	if err {
		return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  true,
		})
	}
	return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
		Data:    nil,
		Message: "err()",
		Status:  false,
	})
}

func (u *UserStructC) GetUsersController(c echo.Context) error {
	users, check := u.userS.GetUsersService()

	if check != nil {
		return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
		Data:    users,
		Message: "Successfull get users account",
		Status:  true,
	})
}

func (u *UserStructC) GetUserController(c echo.Context) error {
	id := c.Param("id")
	user, check := u.userS.GetUserService(id)

	if check != nil {
		return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
		Data:    user,
		Message: "Successfull get user account",
		Status:  true,
	})
}

func (u *UserStructC) CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	_, check := u.userS.CreateService(&user)

	if check != nil {
		return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
		Data:    user,
		Message: "Successfull create users account",
		Status:  true,
	})
}

func (u *UserStructC) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}
	c.Bind(&user)

	dataUser, check := u.userS.UpdateService(&user, id)

	if check != nil {
		return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
		Data:    dataUser,
		Message: "Successfull update user account",
		Status:  true,
	})
}

func (u *UserStructC) DeleteUserController(c echo.Context) error {
	id := c.Param("id")

	check := u.userS.DeleteService(id)

	if check != nil {
		return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
		Data:    id,
		Message: "Successfull delete user account",
		Status:  true,
	})
}
