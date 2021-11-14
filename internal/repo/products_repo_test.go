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

func TestGetAllProductsFromDB(t *testing.T) {
	pgDb, mck := util.SetupMockDB()
	defer pgDb.Close()
	logger := zaptest.NewLogger(t).Sugar()

	t.Run("Get products", func(t *testing.T) {
		expected := []entity.Discount{
			{
				ID: 1,
				Field:      "category",
				Value:      "boots",
				Amount:     30,
				Priority:   1,
				Created_At: time.Now(),
				Updated_At: time.Now(),
			},
			{
				ID: 2,
				Field:      "sku",
				Value:      "000003",
				Amount:     15,
				Priority:   2,
				Created_At: time.Now(),
				Updated_At: time.Now(),
			},
		}

		p := NewDiscountRepository(logger, pgDb)

		mck.
			ExpectQuery(
				regexp.QuoteMeta(`SELECT * FROM discount`)).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "field", "value", "amount", "priority", "created_at", "updated_at"}).
					AddRow(1, "category", "boots", 30, 1, time.Now(), time.Now()).
					AddRow(2, "sku", "000003", 15, 2, time.Now(), time.Now()),
			)

		result, err := p.GetDiscounts(context.Background())

		assert.EqualValues(t, expected, result)
		assert.Nil(t, err)
	})
}
