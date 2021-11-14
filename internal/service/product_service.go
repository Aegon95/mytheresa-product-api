package service

import (
	"context"
	"fmt"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/Aegon95/mytheresa-product-api/internal/model"
	"github.com/Aegon95/mytheresa-product-api/internal/repo"
	"go.uber.org/zap"
)

type ProductService interface {
	GetProducts(ctx context.Context, category entity.Category, than int64) ([]model.ProductDTO, error)
}

type productService struct {
	log *zap.SugaredLogger
	productRepo repo.ProductRepository
	discService DiscountService
}

func NewProductService(log *zap.SugaredLogger, productRepo repo.ProductRepository, discService DiscountService) ProductService {
	return &productService{
		log,
		productRepo,
		discService,
	}
}

func (p *productService) GetProducts(ctx context.Context, category entity.Category, priceLessThan int64) ([]model.ProductDTO, error) {
	products, err := p.productRepo.GetProducts(ctx, category, priceLessThan)
	if err != nil {
		p.log.Errorf("Error occurred while fetching products %v",err)
		return nil, err
	}
	productsDTO := make([]model.ProductDTO, 0, len(products))
	discounts, err := p.discService.GetDiscounts(ctx)
	if err != nil {
		p.log.Errorf("Error occurred while fetching discounts %v",err)
		return nil, err
	}
	for _, product := range products {
		discount := getDiscount(discounts, product.ProductToDiscountRequest())
		originalPrice := product.Price
		finalPrice := originalPrice - (discount*originalPrice)/100
		var discountPercentage string
		if discount > 0 {
			discountPercentage = fmt.Sprintf("%d%%", discount)
		} else {
			discountPercentage = "null"
		}

		productDto := model.ProductDTO{
			SKU:      product.SKU,
			Name:     product.Name,
			Category: product.Category,
			Price: model.Price{
				Original:           originalPrice,
				Final:              finalPrice,
				DiscountPercentage: discountPercentage,
				Currency:           "EUR",
			},
		}
		productsDTO = append(productsDTO, productDto)
	}

	return productsDTO, nil
}

func getDiscount(discounts []entity.Discount, discReq entity.DiscountRequest) int64 {
	dis := entity.DiscountResult{}
	for _, discount := range discounts{
		discRes := discount.CalculateDiscount(discReq)

		if discRes.Priority != 0 && discRes.Priority > dis.Priority{
			dis = discRes
		}
	}

	return dis.Value
}
