package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/kakaw2016/Gotraining/ExtensionSearch/secondprojecttestarea/projectfileslocation"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func IndexFunc(w http.ResponseWriter, r *http.Request) {
	fileaddress := projectfileslocation.Data.Directorypath
	t, err := template.ParseFiles("templates/indexPage.html")
	checkError(err)
	t.Execute(w, fileaddress)
}

func Nofiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/nofiles.html")

}

func Noextension(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/noextension.html")
}

func Filesindirectory(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/filesindirectory.html")

}

func Localsearchengine(w http.ResponseWriter, r *http.Request) {
	p := projectfileslocation.Searchfilepage{
		Title:         "Your favorate interface to locate your files",
		Directorypath: []string{},
		Extensionfind: []string{},
		Filelocation:  "Search results",
	}
	t, _ := template.ParseFiles("/templates/localsearchengine.html")
	t.Execute(w, p)
}
