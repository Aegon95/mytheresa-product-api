package db

import (
	"fmt"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func SeedDatabase(count int, db *sqlx.DB) error {
	err := seedProducts(count, db)

	if err != nil {
		return err
	}
	err = seedDiscounts(db)
	if err != nil {
		return err
	}
	return nil
}

func seedProducts(count int, db *sqlx.DB) error {
	var p entity.Product
	err := db.Get(&p, "SELECT id FROM product LIMIT 1")
	if err != nil {
		products := make([]entity.Product, 0, count)
		for i := 0; i < count; i++ {
			prod := entity.Product{
				SKU:        fmt.Sprintf("%.6d", i+1),
				Name:       RandStringRunes(9),
				Category:   randomCategory(),
				Price:      rand.Int63n(89999) + 10000,
				Created_At: time.Now(),
			}
			products = append(products, prod)
		}
		_, err = db.NamedExec(`INSERT INTO product (sku, name, price, category, created_at)
        VALUES (:sku, :name, :price, :category, :created_at)`, products)
		if err != nil {
			return err
		}

	}
	return nil
}

func seedDiscounts(db *sqlx.DB) error {
	var d entity.Discount
	err := db.Get(&d, "SELECT id FROM discount LIMIT 1")
	if err != nil {
		discounts := []entity.Discount{
			{
				Field:      "category",
				Value:      "boots",
				Amount:     30,
				Priority:   1,
				Created_At: time.Now(),
				Updated_At: time.Now(),
			},
			{
				Field:      "sku",
				Value:      "000003",
				Amount:     15,
				Priority:   2,
				Created_At: time.Now(),
				Updated_At: time.Now(),
			},

		}

		_, err = db.NamedExec(`INSERT INTO discount (field, value, amount, priority, created_at, updated_at)
        VALUES (:field, :value, :amount, :priority, :created_at, :updated_at)`, discounts)
		if err != nil {
			return err
		}

	}
	return nil
}

func randomCategory() entity.Category {
	return entity.CategoryName(rand.Intn(3))
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
