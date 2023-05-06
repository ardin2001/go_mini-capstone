package controllers

import (
	"net/http"
	"strconv"

	"github.com/ardin2001/go_mini-capstone/configs"
	"github.com/ardin2001/go_mini-capstone/helpers"
	"github.com/ardin2001/go_mini-capstone/middlewares"
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

func (u *UserStructC) GetUsersController(c echo.Context) error {
	data, err := middlewares.AdminVerification(c)
	if !err {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: data.Error(),
			Status:  false,
		})
	}

	users, check := u.userS.GetUsersService()
	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    users,
		Message: "Successfull get users account",
		Status:  true,
	})
}

func (u *UserStructC) GetUserController(c echo.Context) error {
	getDataUser := middlewares.GetDataJWT(c)
	id := strconv.Itoa(int(getDataUser.ID))
	user, check := u.userS.GetUserService(id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
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
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    user,
		Message: "Successfull create users account",
		Status:  true,
	})
}

func (u *UserStructC) UpdateUserController(c echo.Context) error {
	getDataUser := middlewares.GetDataJWT(c)
	id := strconv.Itoa(int(getDataUser.ID))
	user := models.User{}
	c.Bind(&user)

	dataUser, check := u.userS.UpdateService(&user, id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    dataUser,
		Message: "Successfull update user account",
		Status:  true,
	})
}

func (u *UserStructC) DeleteUserController(c echo.Context) error {
	getDataUser := middlewares.GetDataJWT(c)
	id := strconv.Itoa(int(getDataUser.ID))

	check := u.userS.DeleteService(id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    id,
		Message: "Successfull delete user account",
		Status:  true,
	})
}

func (us *UserStructC) LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	DB, _ := configs.InitDB()
	err := DB.Where("nama = ? AND password = ?", user.Nama, user.Password).First(&user).Error

	if err != nil {
		return helpers.Response(c, http.StatusUnauthorized, helpers.ResponseModel{
			Data:    nil,
			Message: "login failed username or password",
			Status:  false,
		})
	}

	token, _ := middlewares.CreateToken(user.ID, user.Nama, user.Role)
	userresponse := models.UserResponse{ID: user.ID, Nama: user.Nama, Email: user.Email, Role: user.Role, Token: token}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    userresponse,
		Message: "Login successfull",
		Status:  true,
	})
}
