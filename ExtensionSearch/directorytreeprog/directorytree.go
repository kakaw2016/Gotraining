package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	//directory := "."

	directory := "/home/ybf/Documents/"
	customVisit := func(path string, fileInfo os.FileInfo, err1 error) error {
		if err1 != nil {
			return err1
		}
		fmt.Println(path, fileInfo.Size())
		return nil
	}
	error := filepath.Walk(directory, customVisit)
	if error != nil {
		log.Println(error)
	}

}
