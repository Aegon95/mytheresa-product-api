package repo

import (
	"context"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ProductRepository interface {
	GetProducts(context.Context, entity.Category, int64) ([]entity.Product, error)
}

type productsRepository struct {
	log *zap.SugaredLogger
	db *sqlx.DB
}

func NewProductRepository(log *zap.SugaredLogger, db *sqlx.DB) ProductRepository {
	return &productsRepository{
		log,
		db,
	}
}

func (p *productsRepository) GetProducts(ctx context.Context, category entity.Category, priceLessThan int64) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if category == "" && priceLessThan <= 0 {
		err = p.db.SelectContext(ctx, &products, "SELECT * FROM product LIMIT 5")
	} else if category != "" && priceLessThan <= 0 {
		err = p.db.SelectContext(ctx, &products, "SELECT * FROM product WHERE category = $1 LIMIT 5", category)
	} else if category != "" && priceLessThan > 0 {
		err = p.db.SelectContext(ctx, &products, "SELECT * FROM product WHERE category = $1 AND price <= $2 LIMIT 5", category, priceLessThan)
	} else if category == "" && priceLessThan > 0 {
		err = p.db.SelectContext(ctx, &products, "SELECT * FROM product WHERE price <= $1 LIMIT 5", priceLessThan)
	}

	if err != nil {
		p.log.Errorf("error occurred while fetching records %v", err)
		return nil, err
	}
	return products, err
}
