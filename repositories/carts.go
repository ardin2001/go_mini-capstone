package repositories

import (
	"errors"
	"fmt"

	"github.com/ardin2001/go_mini-capstone/models"
	"gorm.io/gorm"
)

type CartInterfaceR interface {
	GetCartsRepository(id string) ([]models.Cart, error)
	GetCartRepository(id, user_id string) (*models.Cart, error)
	CreateCartRepository(Cart *models.Cart) (*models.Cart, error)
	DeleteCartRepository(id, id_user string) error
	UpdateCartRepository(CartId *models.Cart, id string) (*models.Cart, error)
}

type CartStructR struct {
	DB *gorm.DB
}

func NewCartRepositories(db *gorm.DB) CartInterfaceR {
	return &CartStructR{
		DB: db,
	}
}

func (cr *CartStructR) GetCartsRepository(id string) ([]models.Cart, error) {
	var carts []models.Cart
	var check error
	if id == "" {
		check = cr.DB.Preload("User").Preload("Product").Find(&carts).Error
	} else {
		check = cr.DB.Where("user_id", id).Preload("User").Preload("Product").Find(&carts).Error
	}
	if check != nil {
		return nil, check
	}

	return carts, check
}

func (cr *CartStructR) GetCartRepository(id, user_id string) (*models.Cart, error) {
	var cart models.Cart
	check := cr.DB.Where("user_id", user_id).Preload("User").Preload("Product").First(&cart, id).Error
	if check != nil {
		return nil, check
	}
	return &cart, check
}

func (cr *CartStructR) CreateCartRepository(cart *models.Cart) (*models.Cart, error) {
	check := cr.DB.Save(cart).Error
	if check != nil {
		return nil, check
	}
	return cart, check
}

func (cr *CartStructR) DeleteCartRepository(id, user_id string) error {
	if err := cr.DB.Where("id = ?", id).Take(&models.Cart{}).Error; err != nil {
		return errors.New("not_found")
	}

	check := cr.DB.Where("user_id", user_id).Delete(&models.Cart{}, &id).Error
	fmt.Println(check, user_id)
	if check != nil {
		return errors.New("protected")
	}
	return check
}

func (cr *CartStructR) UpdateCartRepository(cartId *models.Cart, id string) (*models.Cart, error) {
	check := cr.DB.Save(cartId).Error
	if check != nil {
		return nil, check
	}
	return cartId, check
}
