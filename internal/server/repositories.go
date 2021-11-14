package server

import (
	"github.com/Aegon95/mytheresa-product-api/internal/repo"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repositories struct {
	ProductRepository repo.ProductRepository
	DiscountRepository repo.DiscountRepository
}

func SetupRepositories(log *zap.SugaredLogger, db *sqlx.DB) *Repositories {
	productRepository := repo.NewProductRepository(log, db)
	discountRepository := repo.NewDiscountRepository(log, db)
	return &Repositories{
		ProductRepository: productRepository,
		DiscountRepository: discountRepository,
	}
}