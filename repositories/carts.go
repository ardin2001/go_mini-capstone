package repositories

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"gorm.io/gorm"
)

type CartInterfaceR interface {
	GetCartsRepository() ([]models.Cart, error)
	GetCartRepository(id string) (*models.Cart, error)
	CreateCartRepository(Cart *models.Cart) (*models.Cart, error)
	DeleteCartRepository(id string) error
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

func (cr *CartStructR) GetCartsRepository() ([]models.Cart, error) {
	var carts []models.Cart
	check := cr.DB.Preload("User").Preload("Product").Find(&carts).Error

	if check != nil {
		return nil, check
	}

	return carts, check
}

func (cr *CartStructR) GetCartRepository(id string) (*models.Cart, error) {
	var cart models.Cart
	check := cr.DB.First(&cart, id).Error
	if check != nil {
		return nil, check
	}
	return &cart, check
}

func (cr *CartStructR) DeleteCartRepository(id string) error {
	check := cr.DB.Delete(&models.Cart{}, &id).Error
	return check
}

func (cr *CartStructR) CreateCartRepository(cart *models.Cart) (*models.Cart, error) {
	check := cr.DB.Save(cart).Error
	if check != nil {
		return nil, check
	}
	return cart, check
}

func (cr *CartStructR) UpdateCartRepository(cartId *models.Cart, id string) (*models.Cart, error) {
	check := cr.DB.Save(cartId).Error
	if check != nil {
		return nil, check
	}
	return cartId, check
}
