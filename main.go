package main

import (
	"fmt"
	"net/http"
)

const PortNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Iniciando aplicacion en el puerto %s", PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
