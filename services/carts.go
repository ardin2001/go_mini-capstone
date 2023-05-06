package services

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"
)

type CartInterfaceS interface {
	GetCartsService(id string) ([]models.Cart, error)
	GetCartService(id, user_id string) (*models.Cart, error)
	CreateCartService(Cart *models.Cart) (*models.Cart, error)
	UpdateCartService(CartId *models.Cart, id, user_id string) (*models.Cart, error)
	DeleteCartService(id, user_id string) error
	DeleteBatchService(user_id string, carts *[]models.Cart) error
}

type CartStructS struct {
	cartR repositories.CartInterfaceR
}

func NewCartServices(cartR repositories.CartInterfaceR) CartInterfaceS {
	return &CartStructS{
		cartR: cartR,
	}
}

func (cs *CartStructS) GetCartsService(id string) ([]models.Cart, error) {
	carts, err := cs.cartR.GetCartsRepository(id)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (cs *CartStructS) GetCartService(id, user_id string) (*models.Cart, error) {
	cart, err := cs.cartR.GetCartRepository(id, user_id)
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

func (cs *CartStructS) UpdateCartService(cartId *models.Cart, id, user_id string) (*models.Cart, error) {
	getCartId, err := cs.cartR.GetCartRepository(id, user_id)

	if err != nil {
		return nil, err
	}

	if cartId.Jumlah != 0 {
		getCartId.Jumlah = cartId.Jumlah
	}

	cart, err := cs.cartR.UpdateCartRepository(getCartId, id)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (cs *CartStructS) DeleteCartService(id, user_id string) error {
	err := cs.cartR.DeleteCartRepository(id, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CartStructS) DeleteBatchService(user_id string, carts *[]models.Cart) error {
	err := cs.cartR.DeleteBatchRepository(user_id, carts)
	if err != nil {
		return err
	}

	return nil
}
