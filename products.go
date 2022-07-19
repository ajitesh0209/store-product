package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"store-product/models"
	"strconv"
)

func (a *App) createProduct(responseWrite http.ResponseWriter, request *http.Request) {
	var product *models.Product
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&product); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
		respondWithError(responseWrite, http.StatusBadRequest, "Invalid message ")
		return
	}

	defer request.Body.Close()

	if err := product.CreateProduct(a.DB); err != nil {
		respondWithError(responseWrite, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(responseWrite, http.StatusCreated, product)

}

func (a *App) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	count, _ := strconv.Atoi(request.FormValue("count"))
	start, _ := strconv.Atoi(request.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getProducts(a.DB, start, count)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(responseWriter, http.StatusOK, products)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getProducts(db *sql.DB, start, count int) ([]*models.Product, error) {
	rows, err := db.Query("SELECT id, store_id, product_id, is_available FROM store_products limit $1 offset $2", count, start)

	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
		return nil, err
	}

	var products []*models.Product

	defer rows.Close()

	for rows.Next() {
		p := new(models.Product)
		if err := rows.Scan(&p.Id, &p.StoreId, &p.ProductId, &p.IsAvailable); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	rows.Close()

	return products, nil
}
