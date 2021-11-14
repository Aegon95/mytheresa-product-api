package util

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func SetupMockDB() (*sqlx.DB, sqlmock.Sqlmock) {
	database, mck, err := sqlmock.New()
	if err != nil {
		fmt.Printf("can't create sqlmock: %s", err)
	}

	pgDb := sqlx.NewDb(database, "postgres")

	return pgDb, mck
}
