package db

import (
	"fmt"
	"github.com/Aegon95/mytheresa-product-api/config"
	constants "github.com/Aegon95/mytheresa-product-api/internal/util"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Database interface {
	Setup() *sqlx.DB
}

type postgresql struct {
	log *zap.SugaredLogger
	cfg *config.Configuration
}

func NewDatabase(log *zap.SugaredLogger, cfg *config.Configuration) Database {
	return &postgresql{
		log,
		cfg,
	}
}

func (p *postgresql) Setup() *sqlx.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.cfg.Database.Host, p.cfg.Database.Port, p.cfg.Database.Username, p.cfg.Database.Password, p.cfg.Database.Dbname)

	db, err := sqlx.Connect(p.cfg.Database.Driver, psqlInfo)
	if err != nil {
		p.log.Fatalf("Error occurred while connecting to database %v", err)
	}

	db.SetMaxIdleConns(p.cfg.Database.MaxIdleConns)
	db.SetMaxOpenConns(p.cfg.Database.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(p.cfg.Database.MaxLifetime))

	err = db.Ping()

	if err != nil {
		p.log.Fatalf("Error occurred while pinging to database %v", err)
	}

	p.log.Info("Successfully connected to postgres database")

	err = p.runMigration(db)

	if err != nil {
		p.log.Fatalf("Error occurred while running migrations on database %v", err)
	}

	p.log.Info("Starting Database seeding")
	err = SeedDatabase(10000, db)

	if err != nil {
		p.log.Fatalf("Error occurred while seeding data to database %v", err)
	}else{
		p.log.Info("Finished Database seeding")
	}

	return db
}

func (p *postgresql) runMigration(db *sqlx.DB) error {
	p.log.Info("Starting Migrations")
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	src, err := httpfs.New(http.Dir(constants.MIGRATIONS_PATH), "./")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("httpfs", src, p.cfg.Database.Driver, driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	p.log.Info("Finished Migrations")
	return nil
}
