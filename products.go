package main

import (
	"encoding/json"
	"log"
	"net/http"
	"store-product/models"
	"strconv"
)

func (a *App) createProduct(responseWrite http.ResponseWriter, request *http.Request) {
	var product *models.StoreProduct
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(responseWrite, http.StatusBadRequest, "Invalid message ")
		return
	}

	defer request.Body.Close()

	product.CreateProduct(a.DB)

	respondWithJSON(responseWrite, http.StatusCreated, product)

}

func (a *App) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	count, _ := strconv.Atoi(request.FormValue("count"))
	start, _ := strconv.Atoi(request.FormValue("start"))
	var product *models.StoreProduct

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := product.GetProducts(a.DB, start, count)
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
	_, err := w.Write(response)
	if err != nil {
		log.Fatal("Error writing response to JSON", err.Error())
	}
}
