package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/kakaw2016/Gotraining/ExtensionSearch/secondproject/projectfileslocation.go"
)

type Searchfilepage struct {
	Title         string
	Directorypath string
	Extensionfind string
	filelocation  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func localsearchengine(w http.ResponseWriter, r *http.Request) {
	p := Searchfilepage{Title: "Your favorate interface to locate your files", Directorypath: "Path to the directory", Extensionfind: "Array of Extensions", filelocation: "Search results"}
	t, _ := template.ParseFiles("firsttemplate.html")
	t.Execute(w, p)
}

func main() {
	projectfileslocation.Directoryfileslocation()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/localidentification/", localsearchengine)
	http.ListenAndServe(":8000", nil)
}
