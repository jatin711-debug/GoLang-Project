package main

import (
	"fmt"
	"net/http"
	"go-project/pkg/handlers"
)

const portNumber string = ":8080"


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	_ = http.ListenAndServe(portNumber,nil);
}