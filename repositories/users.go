package repositories

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"gorm.io/gorm"
)

type UserInterfaceR interface {
	GetUsersRepository() ([]models.User, error)
	GetUserRepository(id string) (*models.User, error)
	CreateRepository(User *models.User) (*models.User, error)
	DeleteRepository(id string) error
	UpdateRepository(userId *models.User, id string) (*models.User, error)
}

type UserstructR struct {
	DB *gorm.DB
}

func NewUserRepositories(db *gorm.DB) UserInterfaceR {
	return &UserstructR{
		DB: db,
	}
}

func (us *UserstructR) GetUsersRepository() ([]models.User, error) {
	user := models.User{
		Name:     "ardin",
		Email:    "ardin@gmail.com",
		Password: "27sb2d73b",
	}
	arr_users := []models.User{user, user, user}

	err := true
	if err {
		return arr_users, nil
	}

	return nil, nil
}

func (us *UserstructR) GetUserRepository(id string) (*models.User, error) {
	user := models.User{
		Name:     "ardin",
		Email:    "ardin@gmail.com",
		Password: "27sb2d73b",
	}
	err := true
	if err {
		return &user, nil
	}

	return nil, nil
}

func (us *UserstructR) DeleteRepository(id string) error {
	err := true
	if err {
		return nil
	}

	return nil
}

func (us *UserstructR) CreateRepository(user *models.User) (*models.User, error) {
	err := true
	if err {
		return user, nil
	}

	return nil, nil
}

func (us *UserstructR) UpdateRepository(userId *models.User, id string) (*models.User, error) {
	user := models.User{
		Name:     "ardin",
		Email:    "ardin@gmail.com",
		Password: "27sb2d73b",
	}
	err := true
	if err {
		return &user, nil
	}

	return nil, nil
}
