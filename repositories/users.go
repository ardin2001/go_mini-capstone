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
	var users []models.User
	check := us.DB.Find(&users).Error

	if check != nil {
		return nil, check
	}

	return users, check
}

func (us *UserstructR) GetUserRepository(id string) (*models.User, error) {
	var user models.User
	check := us.DB.First(&user, id).Error
	if check != nil {
		return nil, check
	}
	return &user, check
}

func (us *UserstructR) DeleteRepository(id string) error {
	check := us.DB.Delete(&models.User{}, &id).Error
	return check
}

func (us *UserstructR) CreateRepository(user *models.User) (*models.User, error) {
	check := us.DB.Save(user).Error
	if check != nil {
		return nil, check
	}
	return user, check
}

func (us *UserstructR) UpdateRepository(userId *models.User, id string) (*models.User, error) {
	check := us.DB.Save(userId).Error
	if check != nil {
		return nil, check
	}
	return userId, check
}
