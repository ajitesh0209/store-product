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

var storeProductRepository *repository.StoreProductRepository

type IStoreProduct interface {
	AddStoreProduct(writer http.ResponseWriter, request *http.Request)
	GetStoreProductDetails(writer http.ResponseWriter, request *http.Request)
}

func AddStoreProduct(writer http.ResponseWriter, request *http.Request) {
	var products *models.StoreProduct
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&products); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, "Invalid Request Payload | Error :: "+err.Error())
		return
	}

	if err := storeProductRepository.CreateStoreProduct(products); err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(writer, http.StatusCreated, products)
}

func GetStoreProductDetails(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	products, err := storeProductRepository.GetProductForStore(id)
	if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(writer, http.StatusOK, products)

}
