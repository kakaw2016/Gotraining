package projectfileslocation

import (
	"fmt"

	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Searchfilepage struct {
	Extension     string
	Directorypath []string
	Extensionfind []string
	Filelocation  string
}

var Data = Searchfilepage{

	Extension:     "",
	Directorypath: []string{},
	Extensionfind: []string{},
	Filelocation:  "",
}

//var Metadata []Searchfilepage

func removeduplicatestring(datacollected []string) []string {
	alldatacollected := make(map[string]bool)
	correctlist := []string{}
	for _, data := range datacollected {
		if _, value := alldatacollected[data]; !value {
			alldatacollected[data] = true
			correctlist = append(correctlist, data)
		}
	}
	return correctlist
}

func Extensionscollection(path string) []string {

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		Data.Extensionfind = append(Data.Extensionfind, filepath.Ext(path))
		Data.Extensionfind = removeduplicatestring(Data.Extensionfind)

		return nil
	})

	return Data.Extensionfind
}

func Filesfilteredbyextension(root, pattern string) ([]string, error) {

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			Data.Directorypath = append(Data.Directorypath, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return Data.Directorypath, nil
}

/*func Directoryfileslocation() []string {

	fmt.Print("ENTER YOUR DIRECTORY LOCATION: ")

	fmt.Scanln(&Data.Filelocation)

	fmt.Println("First elements", Extensionscollection(Data.Filelocation))

	fmt.Print("SELECT ONE EXTENSION TO BROWSER PATH OF FILES: ")

	fmt.Scanln(&Data.Extension)

	extension := "*." + Data.Extension

	files, err := Filesfilteredbyextension(Data.Filelocation, extension)
	if err != nil {
		fmt.Println("There is a terrible Error", err)
	} else if files == nil {

		fmt.Println("No result from the search")

	} else {

		Data.Directorypath = files
	}

	return Data.Directorypath

}*/

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/*func Filesindirectory(w http.ResponseWriter, r *http.Request) {

	directoryaddress := Data.Filelocation
	t, err := template.ParseFiles("templates/filesindirectory.html")
	checkError(err)
	t.Execute(w, directoryaddress)

}*/

func Filesindirectory(w http.ResponseWriter, r *http.Request) {

	extensioncollect := Extensionscollection(Data.Filelocation)

	t, err := template.ParseFiles("templates/filesindirectory.html")
	checkError(err)
	t.Execute(w, extensioncollect)
}

func Extensioncollector(w http.ResponseWriter, r *http.Request) {

	extensioncollected := Extensionscollection(Data.Filelocation)

	t, err := template.ParseFiles("templates/extensioncollector.html")
	checkError(err)
	t.Execute(w, extensioncollected)
}

func Localsearchengine(w http.ResponseWriter, r *http.Request) {
	localengine := Data.Extension

	t, err := template.ParseFiles("/templates/localsearchengine.html")
	checkError(err)
	t.Execute(w, localengine)
}

func Noextension(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/noextension.html")
}

func IndexFunc(w http.ResponseWriter, r *http.Request) {
	fileaddress, err := Filesfilteredbyextension(Data.Filelocation, `"*." + Data.Extension`)
	checkError(err)

	t, err := template.ParseFiles("templates/indexPage.html")
	checkError(err)
	t.Execute(w, fileaddress)
}

func Nofiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/nofiles.html")

}
