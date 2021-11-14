package service

import (
	"context"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/Aegon95/mytheresa-product-api/internal/repo"
	"go.uber.org/zap"
)

type DiscountService interface {
	GetDiscounts(ctx context.Context) ([]entity.Discount, error)
}

type discountService struct {
	log *zap.SugaredLogger
	discountRepo repo.DiscountRepository
}

func NewDiscountService(log *zap.SugaredLogger, discountRepo repo.DiscountRepository) DiscountService {
	return &discountService{
		log,
		discountRepo,
	}
}

func (p *discountService) GetDiscounts(ctx context.Context) ([]entity.Discount, error) {
	discounts, err := p.discountRepo.GetDiscounts(ctx)
	if err != nil {
		p.log.Errorf("Error occurred while fetching discounts %v",err)
		return nil, err
	}

	return discounts, nil
}
