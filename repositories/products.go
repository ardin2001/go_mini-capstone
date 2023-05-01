package repositories

import (
	"github.com/ardin2001/go_mini-capstone/models"
	"gorm.io/gorm"
)

type ProductInterfaceR interface {
	GetProductsRepository() ([]models.Product, error)
	GetProductRepository(id string) (*models.Product, error)
	CreateProductRepository(Product *models.Product) (*models.Product, error)
	DeleteProductRepository(id string) error
	UpdateProductRepository(ProductId *models.Product, id string) (*models.Product, error)
}

type ProductStructR struct {
	DB *gorm.DB
}

func NewProductRepositories(db *gorm.DB) ProductInterfaceR {
	return &ProductStructR{
		DB: db,
	}
}

func (pr *ProductStructR) GetProductsRepository() ([]models.Product, error) {
	var Products []models.Product
	check := pr.DB.Find(&Products).Error

	if check != nil {
		return nil, check
	}

	return Products, check
}

func (pr *ProductStructR) GetProductRepository(id string) (*models.Product, error) {
	var product models.Product
	check := pr.DB.First(&product, id).Error
	if check != nil {
		return nil, check
	}
	return &product, check
}

func (pr *ProductStructR) DeleteProductRepository(id string) error {
	check := pr.DB.Delete(&models.Product{}, &id).Error
	return check
}

func (pr *ProductStructR) CreateProductRepository(user *models.Product) (*models.Product, error) {
	check := pr.DB.Save(user).Error
	if check != nil {
		return nil, check
	}
	return user, check
}

func (pr *ProductStructR) UpdateProductRepository(userId *models.Product, id string) (*models.Product, error) {
	check := pr.DB.Save(userId).Error
	if check != nil {
		return nil, check
	}
	return userId, check
}
