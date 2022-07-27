package test

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"store-product/models"
	"strconv"
)

func ClearTable(db *gorm.DB) {
	db.Unscoped().Delete(&models.Product{})
	db.Unscoped().Delete(&models.StoreProduct{})
}

func executeRequest(req *http.Request, router *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func addStoreProducts(storeId int, productsCount int, db *gorm.DB) {
	if productsCount < 1 {
		productsCount = 1
	}

	for i := 0; i < productsCount; i++ {
		db.Create(&models.Product{Name: strconv.Itoa(i), Price: float64(10.0 * i)})
		db.Create(&models.StoreProduct{StoreId: storeId, ProductId: i + 1, IsAvailable: true})
	}
}
