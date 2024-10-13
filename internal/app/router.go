package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	api "github.com/jenkka/course-rest-api/api/v1"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/v1", func(r chi.Router) {
		HealthRoutes(r)
	})

	return router
}

func HealthRoutes(router chi.Router) {
	router.Get("/health", api.HealthCheckHandler)
}
