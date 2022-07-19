package models

import (
	"database/sql"
	"fmt"
)

type Product struct {
	Id          int  `json:"id"`
	StoreId     int  `json:"store_id"`
	ProductId   int  `json:"product_id"`
	IsAvailable bool `json:"is_available"`
}

func (p *Product) CreateProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO store_products(store_id, product_id, is_available) VALUES($1, $2, $3) RETURNING id",
		p.StoreId, p.ProductId, p.IsAvailable).Scan(&p.Id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
