package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kugatsuno/udemy-modern-web-apps-golang/pkg/config"
	"github.com/kugatsuno/udemy-modern-web-apps-golang/pkg/handlers"
	"github.com/kugatsuno/udemy-modern-web-apps-golang/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
