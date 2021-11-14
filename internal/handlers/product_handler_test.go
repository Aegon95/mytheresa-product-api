package handlers

import (
	"fmt"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/Aegon95/mytheresa-product-api/internal/model"
	"github.com/Aegon95/mytheresa-product-api/mocks"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zaptest"
	"net/http"
	"net/http/httptest"
	"testing"
)

type productHandlerSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the mocked version of the usecase
	service *mocks.ProductService
	// the functionalities we need to test
	handler ProductHandler
	// testing server to be used the handler
	testingServer *httptest.Server
}

func (suite *productHandlerSuite) SetupSuite() {
	logger := zaptest.NewLogger(suite.T())
	// create a mocked version of usecase
	service := new(mocks.ProductService)
	// inject the usecase to be used by handler
	handler := NewProductHandler(logger.Sugar(), service)

	// create default server using gin, then register all endpoints
	router := chi.NewRouter()
	router.Get("/api/v1/products", handler.GetProducts())


	// create and run the testing server
	testingServer := httptest.NewServer(router)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.testingServer = testingServer
	suite.service = service
	suite.handler = handler
}

func (suite *productHandlerSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

func (suite *productHandlerSuite) TestProductGetAll() {

	products := []model.ProductDTO{
		{
			SKU: "000001",
			Name: "XVlBzgbai",
			Category: "Sandals",
			Price: model.Price{
				Original: 49440,
				Final: 49440,
				DiscountPercentage: "null",
				Currency: "EUR",
			},
		},
		{
			SKU: "000007",
			Name: "PEZQleQYh",
			Category: "Sandals",
			Price: model.Price{
				Original: 20161,
				Final: 20161,
				DiscountPercentage: "null",
				Currency: "EUR",
			},
		},
	}

	suite.service.On("GetProducts", mock.Anything, mock.AnythingOfType("entity.Category"), mock.Anything).Return(products, nil)

	e := httpexpect.New(suite.T(), suite.testingServer.URL)
	path := fmt.Sprintf("/api/v1/products")
	e.GET(path).WithQuery("category", entity.Sandals).WithQuery("priceLessThan",15000).Expect().Status(http.StatusOK).JSON().Array().Equal(products)

	// Check remaining expectations
	suite.service.AssertExpectations(suite.T())

}

func TestProductHandler(t *testing.T) {
	suite.Run(t, new(productHandlerSuite))
}
