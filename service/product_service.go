package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"store-product/models"
	"store-product/repository"
	"store-product/utils"
	"strconv"
)

var productRepository *repository.ProductRepository

type IProducts interface {
	GetProductsById(responseWriter http.ResponseWriter, request *http.Request)
	AddProducts(writer http.ResponseWriter, request *http.Request)
}

func GetProductById(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	products, err := productRepository.GetProduct(id)
	if err != nil {
		utils.RespondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(responseWriter, http.StatusOK, products)
}

func AddProducts(writer http.ResponseWriter, request *http.Request) {
	var products *models.Product
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&products); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := productRepository.CreateProduct(products); err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(writer, http.StatusCreated, products)
}
