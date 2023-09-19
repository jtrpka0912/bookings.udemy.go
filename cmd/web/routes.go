package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jtrpka0912/bookings.udemy.go/pkg/config"
	"github.com/jtrpka0912/bookings.udemy.go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// This is temporary
	mux.Get("/rooms/generals-quarters", handlers.Repo.Generals)
	mux.Get("/rooms/majors-suite", handlers.Repo.Majors)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
