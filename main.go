package store_product

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"store-product/database"
	"store-product/repository"
	"store-product/routers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.InitializeDB()
	repository.InitProductRepository()
	repository.InitStoreProductRepository()
	router := routers.InitializeRoutes()

	log.Println("Running on Port 8010")
	log.Fatal(http.ListenAndServe(":8010", router))
}
