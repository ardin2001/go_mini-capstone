package services

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"
)

type UserInterfaceS interface {
	GetUsersService() ([]models.User, error)
	GetUserService(id string) (*models.User, error)
	CreateService(user *models.User) (*models.User, error)
	UpdateService(userId *models.User, id string) (*models.User, error)
	DeleteService(id string) error
}

type UserStructS struct {
	userR repositories.UserInterfaceR
}

func NewUserServices(userR repositories.UserInterfaceR) UserInterfaceS {
	return &UserStructS{
		userR: userR,
	}
}

func (u *UserStructS) GetUsersService() ([]models.User, error) {
	users, err := u.userR.GetUsersRepository()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserStructS) GetUserService(id string) (*models.User, error) {
	user, err := u.userR.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserStructS) CreateService(user *models.User) (*models.User, error) {
	userR, err := u.userR.CreateRepository(user)
	if err != nil {
		return nil, err
	}

	return userR, nil
}

func (u *UserStructS) UpdateService(userId *models.User, id string) (*models.User, error) {
	user, err := u.userR.UpdateRepository(userId, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserStructS) DeleteService(id string) error {
	err := u.userR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
