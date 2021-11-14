package entity

import "time"

type Discount struct {
	ID         uint64    `json:"id"`
	Field      string    `json:"field"`
	Value      string    `json:"value"`
	Amount     int64     `json:"amount"`
	Priority   int64     `json:"priority"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type DiscountRequest map[string]interface{}

type DiscountResult struct {
	Value int64
	Priority int64
}

func (d Discount)CalculateDiscount(discountReq DiscountRequest) DiscountResult{
	val, ok := discountReq[d.Field]
	actualValue := val.(string)
	if !ok {
		return DiscountResult{}
	}

	if actualValue == d.Value{
		return DiscountResult{
			Value: d.Amount,
			Priority: d.Priority,
		}
	}
	return DiscountResult{}
}