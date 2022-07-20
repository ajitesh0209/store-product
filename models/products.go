package models

import (
	"gorm.io/gorm"
)

type StoreProduct struct {
	Id          int  `json:"id" gorm:"primary_key;auto_increment"`
	StoreId     int  `json:"storeId" gorm:"column:store_id"`
	ProductId   int  `json:"productId" gorm:"product_id"`
	IsAvailable bool `json:"isAvailable" gorm:"is_available"`
}

func (p *StoreProduct) CreateProduct(db *gorm.DB) {
	db.Create(&p)
}

func (p *StoreProduct) CreateMultipleProducts(db *gorm.DB, products []StoreProduct) {
	db.Create(&products)
}

func (p *StoreProduct) GetProducts(db *gorm.DB, start, count int) ([]StoreProduct, error) {
	var products []StoreProduct
	db.Find(&products).Offset(start).Limit(count)
	return products, nil
}
