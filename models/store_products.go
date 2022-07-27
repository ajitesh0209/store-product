package models

type StoreProduct struct {
	Id          int  `json:"id" gorm:"primary_key;auto_increment"`
	StoreId     int  `json:"storeId" gorm:"column:store_id"`
	ProductId   int  `json:"productId" gorm:"product_id"`
	IsAvailable bool `json:"isAvailable" gorm:"is_available"`
}
