package services

import (
	"os"

	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"
	"github.com/joho/godotenv"
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
	godotenv.Load()
	role := os.Getenv("ROLE_U")
	user.Role = role
	userR, err := u.userR.CreateRepository(user)
	if err != nil {
		return nil, err
	}

	return userR, nil
}

func (u *UserStructS) UpdateService(userId *models.User, id string) (*models.User, error) {
	getUserId, err := u.userR.GetUserRepository(id)

	if err != nil {
		return nil, err
	}

	if userId.Nama != "" {
		getUserId.Nama = userId.Nama
	}
	if userId.Email != "" {
		getUserId.Email = userId.Email
	}
	if userId.Password != "" {
		getUserId.Password = userId.Password
	}
	if userId.Role != "" {
		getUserId.Role = userId.Role
	}
	if userId.No_HP != "" {
		getUserId.No_HP = userId.No_HP
	}
	if userId.Alamat != "" {
		getUserId.Alamat = userId.Alamat
	}

	user, err := u.userR.UpdateRepository(getUserId, id)
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
