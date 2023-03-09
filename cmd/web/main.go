// run project from base directory using->  go run .\cmd\web\main.go .\cmd\web\middleware.go .\cmd\web\routes.go

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
	fmt.Printf("Starting application on port %s",portNumber)
	srv := &http.Server{
		Addr: portNumber,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}