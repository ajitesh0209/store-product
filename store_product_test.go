package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a = App{}

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	a.DB.Exec(setupQuery)
}

const setupQuery = `CREATE TABLE IF NOT EXISTS store_products (
    id SERIAL,
    store_id INT NOT NULL,
    product_id INT NOT NULL,
    is_available BOOLEAN,
);`

func clearTable() {
	a.DB.Exec("DELETE FROM store_pr")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

func TestEmptyProductListing(t *testing.T) {
	request, error := http.NewRequest("GET", "/products", nil)
	if error != nil {
		t.Fatal(error)
	}
	executedRequest := a.executeRequest(request)

	if executedRequest.Code == http.StatusOK {
		t.Log("Status mapping correct")
	}

}

func TestStoreProductCreation(t *testing.T) {
	var requestBody = []byte(`{
    	"store_id":2,
    	"product_id":11,
    	"is_available": true
	}`)
	request, error := http.NewRequest("GET", "/product", bytes.NewBuffer(requestBody))
	if error != nil {
		log.Println(error.Error())
		t.Fatal(error)
	}
	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(a.getProducts)

	handler.ServeHTTP(requestRecorder, request)

	var responseBody map[string]interface{}
	err := json.Unmarshal(requestRecorder.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}

	if responseBody["store_id"] == "2" {
		t.Log("Correct Response")
	}
}

func (a *App) executeRequest(request *http.Request) *httptest.ResponseRecorder {
	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(a.getProducts)

	handler.ServeHTTP(requestRecorder, request)
	return requestRecorder
}
