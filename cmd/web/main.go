package main

import (
	"fmt"
	"net/http"

	"github.com/lhhernandez-isc/golang-basis/pkg/handlers"
)

const PortNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Iniciando aplicacion en el puerto %s", PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
