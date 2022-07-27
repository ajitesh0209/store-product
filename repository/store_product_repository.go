package repository

import (
	"store-product/database"
	"store-product/models"
)

var storeProductRepository StoreProductRepository

type IStoreProductRepository interface {
	CreateStoreProduct(product *models.Product) error
	GetProductForStore(id int) ([]models.Product, error)
}

type StoreProductRepository struct {
}

func InitStoreProductRepository() {
	storeProductRepository = StoreProductRepository{}
}

func GetStoreProductRepository() *StoreProductRepository {
	return &storeProductRepository
}

func (repo *StoreProductRepository) GetProductForStore(id int) ([]models.Product, error) {
	db := database.GetConnection()
	var products []models.Product
	txn := db.Select("p.id, p.name, p.price").Table("products p").Joins("INNER JOIN store_products sp ON p.id = sp.product_id").Where("sp.store_id = ? AND sp.is_available=?", id, true).Find(&products)
	return products, txn.Error
}

func (repo *StoreProductRepository) CreateStoreProduct(p *models.StoreProduct) error {
	db := database.GetConnection()
	txn := db.Create(&p)
	return txn.Error
}
