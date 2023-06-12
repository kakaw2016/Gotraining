package main

import (
	"net/http"

	"github.com/kakaw2016/Gotraining/ExtensionSearch/secondprojecttestarea/projectfileslocation"
)

func main() {
	http.HandleFunc("/filesindirectory/", projectfileslocation.Filesindirectory)
	http.HandleFunc("/extensioncollector/", projectfileslocation.Extensioncollector)
	http.HandleFunc("/localsearchengine/", projectfileslocation.Localsearchengine)
	http.HandleFunc("/indexPage/", projectfileslocation.IndexFunc)
	http.HandleFunc("/nofiles/", projectfileslocation.Nofiles)
	http.HandleFunc("/noextension/", projectfileslocation.Noextension)
	http.ListenAndServe(":8080", nil)
}
