package main

import (
	"github.com/OusManDiouf/go-booking-app/pkg/config"
	"github.com/OusManDiouf/go-booking-app/pkg/handlers"
	"github.com/OusManDiouf/go-booking-app/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const (
	portNumber = ":8080"
)

var appConfig config.AppConfig
var sessionManager *scs.SessionManager

func main() {

	// Todo: Change this to true when in production
	appConfig.InProduction = false

	// Initialize Session
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Name = "session_id"
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = appConfig.InProduction // set it true in production !
	// Making the session available to the application
	appConfig.SessionManager = sessionManager

	templateCache, errorTc := render.CreateTemplateCache()
	if errorTc != nil {
		log.Fatal(errorTc)
	}

	appConfig.UseCache = false //devMode
	appConfig.TemplateCache = templateCache

	// Make the appConfig available to the render package
	render.NewTemplates(&appConfig)

	// config the repo
	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	log.Printf("Starting Application on port %s", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfig),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
