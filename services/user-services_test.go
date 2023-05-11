package services

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = NewUserServices(userRMock)
)

func TestGetUsersService_Success(t *testing.T) {
	usersMP := []models.User{
		{
			Nama:     "Mamat",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
		{
			Nama:     "Mamat",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
	}

	usersM := []models.User{
		{
			Nama:     "Mamat",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
		{
			Nama:     "Mamat",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
	}

	userRMock.Mock.On("GetUsersRepository").Return(usersMP, nil)
	users, err := userSMock.GetUsersService()

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, usersM[0].Nama, users[0].Nama)
	assert.Equal(t, usersM[0].Password, users[0].Password)
	assert.Equal(t, usersM[0].Email, users[0].Email)
}

func TestGetUsersService_Failure(t *testing.T) {
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = NewUserServices(userRMock)
	userRMock.Mock.On("GetUsersRepository").Return(nil, errors.New("get all users failed"))
	users, err := userSMock.GetUsersService()

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestGetUserService_Success(t *testing.T) {
	user := models.User{
		Nama:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("GetUserRepository", "1").Return(user, nil)
	users, err := userSMock.GetUserService("1")

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, user.Nama, users.Nama)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
}

func TestGetUserService_Failure(t *testing.T) {
	userRMock.Mock.On("GetUserRepository", "3").Return(nil, fmt.Errorf("user not found"))
	user, err := userSMock.GetUserService("3")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestCreateUserService_Success(t *testing.T) {
	user := models.User{
		Nama:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("CreateRepository", &user).Return(user, nil)
	users, err := userSMock.CreateService(&user)

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, user.Nama, users.Nama)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
}

func TestCreateUserService_Failure(t *testing.T) {
	user := models.User{
		Nama:     "Mamat123",
		Email:    "qwe3123@gmail.com",
		Password: "123456321",
	}

	userRMock.Mock.On("CreateRepository", &user).Return(nil, fmt.Errorf("create user failed"))
	users, err := userSMock.CreateService(&user)

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestDeleteUserService_Success(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := userSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteUserService_Failure(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("user not found"))
	err := userSMock.DeleteService("2")

	assert.NotNil(t, err)
}
