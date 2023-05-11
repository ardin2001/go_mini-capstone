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
	productRMock = &repositories.IProductRepositoryMock{Mock: mock.Mock{}}
	productSMock = NewProductServices(productRMock)
)

func TestGetProductsService_Success(t *testing.T) {
	productsMP := []models.Product{
		{
			Nama:      "Mamat",
			Deskripsi: "produk kualitas terbaik",
			Harga:     14000,
		},
		{
			Nama:      "Mamat",
			Deskripsi: "produk kualitas terbaik",
			Harga:     14000,
		},
	}

	productsM := []models.Product{
		{
			Nama:      "Mamat",
			Deskripsi: "produk kualitas terbaik",
			Harga:     14000,
		},
		{
			Nama:      "Mamat",
			Deskripsi: "produk kualitas terbaik",
			Harga:     14000,
		},
	}

	productRMock.Mock.On("GetProductsRepository").Return(productsMP, nil)
	products, err := productSMock.GetProductsService()

	assert.Nil(t, err)
	assert.NotNil(t, products)

	assert.Equal(t, productsM[0].Nama, products[0].Nama)
	assert.Equal(t, productsM[0].Harga, products[0].Harga)
	assert.Equal(t, productsM[0].Deskripsi, products[0].Deskripsi)
}

func TestGetProducsService_Failure(t *testing.T) {
	productRMock = &repositories.IProductRepositoryMock{Mock: mock.Mock{}}
	productSMock = NewProductServices(productRMock)
	productRMock.Mock.On("GetProductsRepository").Return(nil, errors.New("get all producs failed"))
	producs, err := productSMock.GetProductsService()

	assert.Nil(t, producs)
	assert.NotNil(t, err)
}

func TestGetProductService_Success(t *testing.T) {
	product := models.Product{
		Nama:      "Mamat",
		Deskripsi: "produk kualitas terbaik",
		Harga:     14000,
	}

	productRMock.Mock.On("GetProductRepository", "1").Return(product, nil)
	products, err := productSMock.GetProductService("1")

	assert.Nil(t, err)
	assert.NotNil(t, products)

	assert.Equal(t, product.Nama, products.Nama)
	assert.Equal(t, product.Harga, products.Harga)
	assert.Equal(t, product.Deskripsi, products.Deskripsi)
}

func TestGetProductService_Failure(t *testing.T) {
	productRMock.Mock.On("GetProductRepository", "3").Return(nil, fmt.Errorf("product not found"))
	product, err := productSMock.GetProductService("3")

	assert.NotNil(t, err)
	assert.Nil(t, product)
}

func TestCreateProductService_Success(t *testing.T) {
	product := models.Product{
		Nama:      "Mamat",
		Deskripsi: "produk kualitas terbaik",
		Harga:     14000,
	}

	productRMock.Mock.On("CreateProductRepository", &product).Return(product, nil)
	products, err := productSMock.CreateProductService(&product)

	assert.Nil(t, err)
	assert.NotNil(t, products)

	assert.Equal(t, product.Nama, products.Nama)
	assert.Equal(t, product.Harga, products.Harga)
	assert.Equal(t, product.Deskripsi, products.Deskripsi)
}

func TestCreateProductService_Failure(t *testing.T) {
	product := models.Product{
		Nama:      "Mamat123",
		Deskripsi: "produk kualitas terbaik",
		Harga:     140000,
	}

	productRMock.Mock.On("CreateProductRepository", &product).Return(nil, fmt.Errorf("create product failed"))
	products, err := productSMock.CreateProductService(&product)

	assert.Nil(t, products)
	assert.NotNil(t, err)
}

func TestDeleteProductService_Success(t *testing.T) {
	productRMock.Mock.On("DeleteProductRepository", "1").Return(nil)
	err := productSMock.DeleteProductService("1")

	assert.Nil(t, err)
}

func TestDeleteProductService_Failure(t *testing.T) {
	productRMock.Mock.On("DeleteProductRepository", "2").Return(fmt.Errorf("product not found"))
	err := productSMock.DeleteProductService("2")

	assert.NotNil(t, err)
}
