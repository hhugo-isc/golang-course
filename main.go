package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y, f float32
	x = 100.0
	y = 10.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Iniciando aplicacion en el puerto %s", PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
