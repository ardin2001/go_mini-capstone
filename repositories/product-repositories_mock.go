package repositories

import (
	"fmt"

	"github.com/ardin2001/go_mini-capstone/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock interface {
	GetProductsRepository() ([]models.Product, error)
	GetProductRepository(id string) (*models.Product, error)
	CreateProductRepository(productData *models.Product) (*models.Product, error)
	UpdateProductRepository(productBody *models.Product, id string) (*models.Product, error)
	DeleteProductRepository(id string) error
}

type IProductRepositoryMock struct {
	Mock mock.Mock
}

func NewProductRepositoryMock(mock mock.Mock) ProductRepositoryMock {
	return &IProductRepositoryMock{
		Mock: mock,
	}
}

func (u *IProductRepositoryMock) GetProductsRepository() ([]models.Product, error) {
	args := u.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	products := args.Get(0).([]models.Product)

	return products, nil
}

func (u *IProductRepositoryMock) GetProductRepository(id string) (*models.Product, error) {
	args := u.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	product := args.Get(0).(models.Product)
	return &product, nil
}

func (u *IProductRepositoryMock) CreateProductRepository(productData *models.Product) (*models.Product, error) {
	args := u.Mock.Called(productData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	product := args.Get(0).(models.Product)

	return &product, nil
}

func (u *IProductRepositoryMock) UpdateProductRepository(productData *models.Product, id string) (*models.Product, error) {
	args := u.Mock.Called(id, productData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	product := args.Get(0).(models.Product)

	return &product, nil
}

func (u *IProductRepositoryMock) DeleteProductRepository(id string) error {
	args := u.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}
