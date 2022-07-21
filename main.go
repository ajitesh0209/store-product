package main

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

	log.Fatal(http.ListenAndServe(":8010", router))
}
