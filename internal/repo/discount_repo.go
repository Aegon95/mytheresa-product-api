package repo

import (
	"context"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DiscountRepository interface {
	GetDiscounts(context.Context) ([]entity.Discount, error)
}

type discountRepository struct {
	log *zap.SugaredLogger
	db *sqlx.DB
}

func NewDiscountRepository(log *zap.SugaredLogger, db *sqlx.DB) DiscountRepository {
	return &discountRepository{
		log,
		db,
	}
}

func (p *discountRepository) GetDiscounts(ctx context.Context) ([]entity.Discount, error) {
	var discounts []entity.Discount
	err := p.db.SelectContext(ctx, &discounts, "SELECT * FROM discount")
	if err != nil {
		p.log.Errorf("error occurred while fetching records %v", err)
		return nil, err
	}
	return discounts, err
}
