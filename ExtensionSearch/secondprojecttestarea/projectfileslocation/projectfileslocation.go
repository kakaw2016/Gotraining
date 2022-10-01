package projectfileslocation

import (
	"fmt"

	"os"
	"path/filepath"
)

type Searchfilepage struct {
	Title         string
	Extension     string
	Directorypath []string
	Extensionfind []string
	Filelocation  string
}

/*type AllSearchfilepage struct {
	Searchfilepage []*Searchfilepage
}*/

var Data = Searchfilepage{
	"",
	"",
	[]string{},
	[]string{},
	"",
}

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

func Directoryfileslocation() []string {

	fmt.Print("ENTER YOUR DIRECTORY LOCATION: ")

	fmt.Scanln(&Data.Filelocation)

	fmt.Println("First elements", Extensionscollection(Data.Filelocation))

	fmt.Print("SELECT ONE EXTENSION TO BROWSER PATH OF FILES: ")

	//var ext string

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

}
