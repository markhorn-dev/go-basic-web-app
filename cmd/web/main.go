package main

import (
	"go-basic-web-app/pkg/config"
	"go-basic-web-app/pkg/handlers"
	"go-basic-web-app/pkg/render"

	"log"
	"net/http"
)

// main is the main function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	_ = http.ListenAndServe(":8080", nil)
}
