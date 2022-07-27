package repository

import (
	"log"
	"store-product/database"
	"store-product/models"
)

var productRepository ProductRepository

type IProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProduct(id int) (*models.Product, error)
}

type ProductRepository struct {
}

func InitProductRepository() {
	productRepository = ProductRepository{}
}

func GetProductRepository() *ProductRepository {
	return &productRepository
}

func (repo *ProductRepository) GetProduct(id int) (*models.Product, error) {
	db := database.GetConnection()
	p := &models.Product{}
	txn := db.First(&p, id)
	return p, txn.Error
}

func (repo *ProductRepository) CreateProduct(p *models.Product) error {
	db := database.GetConnection()
	txn := db.Create(&p)
	if txn.Error != nil {
		log.Fatalf("Error Creating Product %v", txn.Error)
	}
	return txn.Error
}
