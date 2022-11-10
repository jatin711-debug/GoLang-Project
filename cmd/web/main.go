package main

import (
	"fmt"
	"go-project/pkg/config"
	"go-project/pkg/handlers"
	"go-project/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber string = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	srv := &http.Server{
		Addr: portNumber,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}