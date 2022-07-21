package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var dbConnection *gorm.DB

func InitializeDB() {
	host := os.Getenv("APP_DB_HOST")
	port := os.Getenv("APP_DB_PORT")
	user := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	dbName := os.Getenv("APP_DB_NAME")
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	fmt.Println(connectionString)
	var err error
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dbConnection = db
}

func GetConnection() *gorm.DB {
	return dbConnection
}
