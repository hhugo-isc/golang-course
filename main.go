package main

import (
	"fmt"
	"net/http"
)

const PortNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValues(2, 3)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page, and 2 + 3 is %d", sum))
}

func addValues(x, y int) (int, error) {
	return x + y, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Iniciando aplicacion en el puerto %s", PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
