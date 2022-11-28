package projectfileslocation

import (
	"fmt"

	"os"
	"path/filepath"
)

func removeDuplicateString(datacollected []string) []string {
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

func extensionsCollection(path string) []string {
	var extension []string
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		extension = append(extension, filepath.Ext(path))
		extension = removeDuplicateString(extension)

		return nil
	})

	return extension
}

func filesFilteredbyExtension(root, pattern string) ([]string, error) {
	var matches []string

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
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func Directoryfileslocation() {

	var repertoryfile string
	fmt.Print("ENTER YOUR DIRECTORY LOCATION: ")

	fmt.Scanln(&repertoryfile)

	fmt.Println("First elements", extensionsCollection(repertoryfile))

	fmt.Print("SELECT ONE EXTENSION TO BROWSER PATH OF FILES: ")

	var ext string

	fmt.Scanln(&ext)

	extension := "*." + ext

	files, err := filesFilteredbyExtension(repertoryfile, extension)
	if err != nil {
		fmt.Println("There is a terrible Error", err)
	} else if files == nil {

		fmt.Println("No result from the search")

	} else {
		for _, locationpathfiles := range files {
			fmt.Println(locationpathfiles)

		}

	}

	fmt.Println("Program ends")
}
