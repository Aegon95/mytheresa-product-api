package server

import (
	"github.com/Aegon95/mytheresa-product-api/internal/handlers"
	"go.uber.org/zap"
)

type Handlers struct {
	ProductHandler handlers.ProductHandler
}

func SetupHandlers(log *zap.SugaredLogger, services *Services) *Handlers {
	productHandlers := handlers.NewProductHandler(log, services.ProductService)

	return &Handlers{
		ProductHandler: productHandlers,
	}
}