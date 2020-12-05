package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetRouter() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	setScraperRoutes(mux)

	return mux
}
