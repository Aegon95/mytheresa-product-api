package service

import (
	"context"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/Aegon95/mytheresa-product-api/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zaptest"
	"testing"
	"time"
)

type discountServiceSuite struct {

	suite.Suite

	repository *mocks.DiscountRepository

	service DiscountService
}

func (suite *discountServiceSuite) SetupTest() {

	logger := zaptest.NewLogger(suite.T()).Sugar()
	repository := new(mocks.DiscountRepository)

	service := NewDiscountService(logger, repository)


	suite.repository = repository
	suite.service = service
}

func (suite *discountServiceSuite) TestGetDiscounts_Positive() {


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

	ctx := context.Background()

	suite.repository.On("GetDiscounts", ctx ).Return(discounts, nil)

	result, err := suite.service.GetDiscounts(ctx)
	suite.Nil(err, "no error when return the discounts")
	suite.Equal(discounts, result, "result and discounts should be equal")
}

func TestDiscountService(t *testing.T) {
	suite.Run(t, new(discountServiceSuite))
}
