package render

import (
	"net/http"
	"fmt"
	"html/template"
)
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ :=  template.ParseFiles("./templates/"+ tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing parsed file temp",err)
	}
}