package main

import (
	"fmt"
	"net/http"
	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w,r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)
	crsfHandler.SetBaseCookie(
		http.Cookie{
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure: app.InProduction,
			Path:    "/",
		},
	)
	return crsfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}