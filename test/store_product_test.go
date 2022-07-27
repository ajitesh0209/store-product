package test

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"store-product/database"
	"store-product/models"
	"store-product/repository"
	"store-product/routers"
	"testing"
)

type StoreProductTestSuite struct {
	suite.Suite
	rtr *mux.Router
	db  *gorm.DB
}

func TestStoreProductTestSuite(t *testing.T) {
	suite.Run(t, new(StoreProductTestSuite))
}

func (suite *StoreProductTestSuite) SetupSuite() {
	database.InitializeDB()
	repository.InitProductRepository()
	repository.InitStoreProductRepository()
	suite.rtr = routers.InitializeRoutes()
	suite.db = database.GetConnection()
}

func (suite *StoreProductTestSuite) TestGetStoreProducts() {
	addStoreProducts(1, 1, suite.db)

	req, _ := http.NewRequest("GET", "/store/1/products", nil)
	response := executeRequest(req, suite.rtr)

	suite.Equal(http.StatusOK, response.Code)
}

func (suite *StoreProductTestSuite) TestGetStoreProductsInvalidID() {
	addStoreProducts(1, 1, suite.db)

	req, _ := http.NewRequest("GET", "/store/a_b/products", nil)
	response := executeRequest(req, suite.rtr)

	suite.Equal(http.StatusBadRequest, response.Code)
}

func (suite *StoreProductTestSuite) TestAddStoreProducts() {
	var jsonStr = []byte(`{"name":"test product", "price": 11.22}`)
	req, _ := http.NewRequest("POST", "/store/3/products", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req, suite.rtr)
	suite.Equal(http.StatusCreated, response.Code)

	products := make([]models.Product, 0)
	err := json.Unmarshal(response.Body.Bytes(), &products)
	if err != nil {
		return
	}

	suite.Equal(len(products), 1)
	suite.Equal(products[0].Name, "test product")
	suite.Equal(products[0].Price, 11.22)
	suite.Equal(int(products[0].Id), 1)
}

func (suite *StoreProductTestSuite) TestAddStoreProductInvalidStoreID() {
	var jsonStr = []byte(`[{"name":"test product", "price": 11.22}]`)
	req, _ := http.NewRequest("POST", "/store/a_b", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req, suite.rtr)
	suite.Equal(http.StatusNotFound, response.Code)
}

func (suite *StoreProductTestSuite) TestAddStoreProductNonExistentStoreID() {
	var jsonStr = []byte(`[{"name":"test product", "price": 11.22}]`)
	req, _ := http.NewRequest("POST", "/store/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req, suite.rtr)
	suite.Equal(http.StatusNotFound, response.Code)
}

func (suite *StoreProductTestSuite) TestAddStoreProductInvalidPayload() {
	var jsonStr = []byte(`[{"name":123, "price": 11.22}]`)
	req, _ := http.NewRequest("POST", "/store/2/products", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req, suite.rtr)
	suite.Equal(http.StatusBadRequest, response.Code)
}
