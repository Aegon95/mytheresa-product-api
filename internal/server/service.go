package server

import (
	"github.com/Aegon95/mytheresa-product-api/internal/service"
	"go.uber.org/zap"
)

type Services struct {
	ProductService service.ProductService
	DiscountService service.DiscountService
}

func SetupServices(log *zap.SugaredLogger, repos *Repositories) *Services {
	discountService := service.NewDiscountService(log, repos.DiscountRepository)
	productService := service.NewProductService(log, repos.ProductRepository, discountService)

	return &Services{
		ProductService: productService,
		DiscountService: discountService,
	}
}