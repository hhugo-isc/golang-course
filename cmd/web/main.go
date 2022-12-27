package main

import (
	"fmt"
	"net/http"

	"github.com/lhhernandez-isc/golang-basis/pkg/config"
	"github.com/lhhernandez-isc/golang-basis/pkg/handlers"
	"github.com/lhhernandez-isc/golang-basis/pkg/render"
)

const PortNumber = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Iniciando aplicacion en el puerto %s\n\n", PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
