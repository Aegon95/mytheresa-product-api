package model

import "github.com/Aegon95/mytheresa-product-api/internal/db/entity"

type ProductDTO struct {
	SKU      string          `json:"sku"`
	Name     string          `json:"name"`
	Category entity.Category `json:"category"`
	Price    Price           `json:"price"`
}

type Price struct {
	Original           int64  `json:"original"`
	Final              int64  `json:"final"`
	DiscountPercentage string `json:"discount_percentage"`
	Currency           string `json:"currency"`
}
