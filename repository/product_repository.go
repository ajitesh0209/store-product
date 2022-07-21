package repository

import (
	"gorm.io/gorm"
	"store-product/models"
)

var productRepository ProductRepository

type IProductRepository interface {
	CreateProduct(db *gorm.DB, product *models.Product) error
	GetProduct(db *gorm.DB, id int) (*models.Product, error)
}

type ProductRepository struct {
}

func InitProductRepository() {
	productRepository = ProductRepository{}
}

func GetProductRepository() *ProductRepository {
	return &productRepository
}

func (repo *ProductRepository) GetProduct(db *gorm.DB, id int) (*models.Product, error) {
	p := &models.Product{}
	txn := db.First(&p, id)
	return p, txn.Error
}

func (repo *ProductRepository) CreateProduct(db *gorm.DB, p *models.Product) error {
	txn := db.Create(&p)
	return txn.Error
}
