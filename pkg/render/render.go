package render

import (
	"go-project/pkg/models"
	"go-project/pkg/config"
	"path/filepath"
	"html/template"
	"net/http"
	"bytes"
	"fmt"
	"log"
)

var app *config.AppConfig

var functions = template.FuncMap{}

func NewTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(" Could not get template for template err")
	}
	buf := new(bytes.Buffer)
	td = addDefaultData(td)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Buff err", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	pagesCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return pagesCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return pagesCache, err
		}
		matches, err := filepath.Glob("./templates/*.page.tmpl")
		if err != nil {
			return pagesCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return pagesCache, err
			}
		}
		pagesCache[name] = ts
	}
	return pagesCache, nil
}
