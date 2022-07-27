package routers

import (
	"github.com/gorilla/mux"
	"store-product/service"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/product/{id:[0-9]+}", service.GetProductById).Methods("GET")
	router.HandleFunc("/products", service.AddProducts).Methods("POST")
	router.HandleFunc("/store/{id}/products", service.AddStoreProduct).Methods("POST")
	router.HandleFunc("/store/{id}/products", service.GetStoreProductDetails).Methods("GET")
	return router
}
