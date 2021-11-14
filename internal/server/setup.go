package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func Setup(logger *zap.SugaredLogger, db *sqlx.DB) *chi.Mux {
	r := chi.NewRouter()
	setupMiddlewares(r)
	repos := SetupRepositories(logger, db)
	services := SetupServices(logger, repos)
	hls := SetupHandlers(logger, services)
	SetupRoutes(r, hls)
	return r
}

func setupMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func SetupRoutes(r *chi.Mux, h *Handlers) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/products", h.ProductHandler.GetProducts())
	})

}
