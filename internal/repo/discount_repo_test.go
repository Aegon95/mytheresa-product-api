package repo

import (
	"context"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/Aegon95/mytheresa-product-api/internal/util"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"regexp"
	"testing"
	"time"
)



func TestGetDiscountsFromDB(t *testing.T) {
	pgDb, mck := util.SetupMockDB()
	defer pgDb.Close()
	logger := zaptest.NewLogger(t).Sugar()

	t.Run("Get discounts", func(t *testing.T) {
		expected := []entity.Product{
			{
				ID:         1,
				SKU:        "000006",
				Name:       "bCsNVlgTe",
				Category:   "sandals",
				Price:      52116,
				Created_At: time.Now(),
				Updated_At: time.Now(),
			},
			{
				ID:         2,
				SKU:        "000014",
				Name:       "yiNKAReKJ",
				Category:   "sandals",
				Price:      20932,
				Created_At: time.Now(),
				Updated_At: time.Now(),
			},
		}

		p := NewProductRepository(logger,pgDb)

		mck.
			ExpectQuery(
				regexp.QuoteMeta(`SELECT * FROM product WHERE category = $1 AND price <= $2  LIMIT 5`)).
			WithArgs(entity.Sneakers, 17000).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "sku", "name", "category", "price", "created_at", "updated_at" }).
					AddRow(1, "000006","bCsNVlgTe", "sandals", 52116, time.Now(), time.Now()).
					AddRow(2, "000014","yiNKAReKJ", "sandals", 20932, time.Now(), time.Now()),
			)

		result, err := p.GetProducts(context.Background(), entity.Sneakers, 17000)

		assert.EqualValues(t, expected, result)
		assert.Nil(t, err)
	})
}
