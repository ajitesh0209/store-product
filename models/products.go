package models

type Product struct {
	Id    int     `json:"id" gorm:"primary_key;auto_increment"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
