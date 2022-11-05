package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(w)

	parsedTemplate, _ :=  template.ParseFiles("./templates/"+ tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing parsed file temp",err)
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template,error) {
	pagesCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return pagesCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("pahge is",page)
		ts,err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return pagesCache, err
		}

		matches, err := filepath.Glob("./*.layout.tmpl")
		if err != nil {
			return pagesCache,err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*/layout.tmpl")
			if err != nil {
				return pagesCache, err
			}
		}
		pagesCache[name] = ts
	}
	return pagesCache,nil
}