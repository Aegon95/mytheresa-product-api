package entity

import "time"

type Product struct {
	ID         uint64    `json:"id"`
	SKU        string    `json:"sku"`
	Name       string    `json:"name"`
	Category   Category  `json:"category"`
	Price      int64     `json:"price"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func (p Product) ProductToDiscountRequest() DiscountRequest{
	return DiscountRequest{
		"category": string(p.Category),
		"sku": p.SKU,
	}
}
