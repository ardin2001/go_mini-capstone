package services

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"
)

type CartInterfaceS interface {
	GetCartsService() ([]models.Cart, error)
	GetCartService(id string) (*models.Cart, error)
	CreateCartService(Cart *models.Cart) (*models.Cart, error)
	UpdateCartService(CartId *models.Cart, id string) (*models.Cart, error)
	DeleteCartService(id string) error
}

type CartStructS struct {
	cartR repositories.CartInterfaceR
}

func NewCartServices(cartR repositories.CartInterfaceR) CartInterfaceS {
	return &CartStructS{
		cartR: cartR,
	}
}

func (cs *CartStructS) GetCartsService() ([]models.Cart, error) {
	carts, err := cs.cartR.GetCartsRepository()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (cs *CartStructS) GetCartService(id string) (*models.Cart, error) {
	cart, err := cs.cartR.GetCartRepository(id)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (cs *CartStructS) CreateCartService(cart *models.Cart) (*models.Cart, error) {
	cartR, err := cs.cartR.CreateCartRepository(cart)
	if err != nil {
		return nil, err
	}

	return cartR, nil
}

func (cs *CartStructS) UpdateCartService(cartId *models.Cart, id string) (*models.Cart, error) {
	getCartId, err := cs.cartR.GetCartRepository(id)

	if err != nil {
		return nil, err
	}

	cart, err := cs.cartR.UpdateCartRepository(getCartId, id)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (cs *CartStructS) DeleteCartService(id string) error {
	err := cs.cartR.DeleteCartRepository(id)
	if err != nil {
		return err
	}

	return nil
}
