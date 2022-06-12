package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/isshougai/rental-bookings/pkg/config"
	"github.com/isshougai/rental-bookings/pkg/handlers"
	"github.com/isshougai/rental-bookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var store cookie.Store

// main is the main application function
func main() {
	app.InProduction = false

	store = cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24,
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	app.Store = store

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
