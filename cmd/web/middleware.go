package main

import (
	"github.com/justinas/nosurf"
	"log"
	"net/http"
)

// WriteToConsole Basic Logging Demo
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("----> Page: ", request.RequestURI)
		next.ServeHTTP(writer, request)
	})
}

//NoSurf setting CSRFToken as middleware to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfig.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad load and save session between request to and from the client in a cookie
func SessionLoad(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}
