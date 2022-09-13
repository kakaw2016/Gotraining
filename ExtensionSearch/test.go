package main

import (
	"fmt"
	
	"os"
	"path/filepath"
)

func extensionscollection(path string) []string {
	var extension []string
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		extension= append(extension, filepath.Ext(path))
		/*var s string
		for i:=0; i < len(extension); i++{		
			if s != extension[i]{			
				extension = append(extension, s)
			}
		}*/
				
		return nil
	})
	

	return extension
}

func filesfilteredbyextension(root, pattern string) ([]string, error) {
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


func main() {

	fmt.Println("First elements", extensionscollection("/home/ybf/Documents"))

	/*files, err := filesfilteredbyextension("/home/ybf/Documents", "*.pdf")
	if err != nil {
		fmt.Println("There is a terrible Error", err)
	} else {
		fmt.Println("The search results are : \n",files)
	}*/

	fmt.Println("Program ends")
}
