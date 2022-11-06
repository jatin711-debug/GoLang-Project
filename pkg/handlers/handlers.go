package handlers

import (
	"go-project/pkg/config"
	"go-project/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {	
	render.RenderTemplate(w, "about.page.tmpl")
}