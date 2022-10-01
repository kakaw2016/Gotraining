package main

import (
	"net/http"

	"github.com/kakaw2016/Gotraining/ExtensionSearch/secondproject/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexFunc)

	http.HandleFunc("/nofiles/", handlers.Nofiles)

	http.HandleFunc("/noextension/", handlers.Noextension)

	http.HandleFunc("/filesindirectory/", handlers.Filesindirectory)
	http.HandleFunc("/filesindirectory/localsearchengine", handlers.Localsearchengine)

	http.ListenAndServe(":8080", nil)
}
