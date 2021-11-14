package service

import (
	"context"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/Aegon95/mytheresa-product-api/internal/model"
	"github.com/Aegon95/mytheresa-product-api/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zaptest"
	"testing"
	"time"
)

type productServiceSuite struct {

	suite.Suite

	repository *mocks.ProductRepository

	service ProductService

	discService *mocks.DiscountService
}

func (suite *productServiceSuite) SetupTest() {

	logger := zaptest.NewLogger(suite.T()).Sugar()
	repository := new(mocks.ProductRepository)
	discountServ := new(mocks.DiscountService)
	service := NewProductService(logger, repository, discountServ)


	suite.repository = repository
	suite.service = service
	suite.discService = discountServ
}

func (suite *productServiceSuite) TestGetProducts_Positive() {

	products := []entity.Product{
		{
			ID:         1,
			SKU:        "000006",
			Name:       "bCsNVlgTe",
			Category:   "sandals",
			Price:      52116,
			Created_At: time.Now(),
		},
		{
			ID:         2,
			SKU:        "000014",
			Name:       "yiNKAReKJ",
			Category:   "sandals",
			Price:      20932,
			Created_At: time.Now(),
		},
	}

	discounts := []entity.Discount{
		{
			Field:      "category",
			Value:      "boots",
			Amount:     30,
			Priority:   1,
			Created_At: time.Now(),
			Updated_At: time.Now(),
		},
		{
			Field:      "sku",
			Value:      "000003",
			Amount:     15,
			Priority:   2,
			Created_At: time.Now(),
			Updated_At: time.Now(),
		},

	}

	productsDto := []model.ProductDTO{
		{
			SKU: "000006",
			Name: "bCsNVlgTe",
			Category: "sandals",
			Price: model.Price{
				Original: 52116,
				Final: 52116,
				DiscountPercentage: "null",
				Currency: "EUR",
			},
		},
		{
			SKU: "000014",
			Name: "yiNKAReKJ",
			Category: "sandals",
			Price: model.Price{
				Original: 20932,
				Final: 20932,
				DiscountPercentage: "null",
				Currency: "EUR",
			},
		},
	}

	ctx := context.Background()
	var category entity.Category = entity.Boots
	priceLessThan := int64(15000)
	suite.discService.On("GetDiscounts", ctx).Return(discounts, nil)
	suite.repository.On("GetProducts", ctx,category, priceLessThan ).Return(products, nil)

	result, err := suite.service.GetProducts(ctx, category, priceLessThan)
	suite.Nil(err, "no error when return the products")
	suite.Equal(productsDto, result, "result and productsDto should be equal")
}

func TestProductService(t *testing.T) {
	suite.Run(t, new(productServiceSuite))
}