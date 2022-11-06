package main

import (
	"go-project/pkg/config"
	"go-project/pkg/handlers"
	"net/http"
	"github.com/go-chi/chi"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}