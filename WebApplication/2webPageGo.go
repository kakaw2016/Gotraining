package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {

	sum, _ := Addvalues(3, 3)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 3 + 3 is %d", sum))
}

func Addvalues(x, y int) (int, error) {

	return x + y, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
