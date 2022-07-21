package repository

import (
	"gorm.io/gorm"
	"store-product/models"
)

var storeProductRepository StoreProductRepository

type IStoreProductRepository interface {
	CreateStoreProduct(db *gorm.DB, product *models.Product) error
	GetProductForStore(db *gorm.DB, id int) (*models.Product, error)
}

type StoreProductRepository struct {
}

func InitStoreProductRepository() {
	storeProductRepository = StoreProductRepository{}
}

func GetStoreProductRepository() *StoreProductRepository {
	return &storeProductRepository
}

func (repo *StoreProductRepository) GetProductForStore(db *gorm.DB, id int) ([]models.Product, error) {
	var products []models.Product
	txn := db.Raw("SELECT products.id, products.name, products.price FROM products inner join store_products on products.id = store_products.product_id WHERE store_products.store_id = ? AND store_products.is_available = true", id).Scan(&products)
	return products, txn.Error
}

func (repo *StoreProductRepository) CreateStoreProduct(db *gorm.DB, p *models.StoreProduct) error {
	txn := db.Create(&p)
	return txn.Error
}
