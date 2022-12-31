package configurationFile

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Duplicate removal in an array inside the code

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

// Extensions identification inside a directory of divers length

func extensionsCollection(filepaths string) []string {
	var extension []string
	filepath.Walk(filepaths, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		extension = append(extension, filepath.Ext(path))
		extension = removeDuplicateString(extension)

		return nil
	})

	return extension
}

// Configuration details in file that generates all the input we will use to find the files

func ReadConfigurationFile(configurationFilepath string) map[string]string {

	// Open the config file in read-only mode
	configFile, err := os.Open(configurationFilepath)
	if err != nil {
		fmt.Println(err)

	}
	defer configFile.Close()

	// Create a new scanner to read the config file
	scanner := bufio.NewScanner(configFile)

	data := make(map[string]string)

	// Read the config file line by line
	for scanner.Scan() {
		// Get the current line
		line := scanner.Text()

		allExtensions := extensionsCollection(line)

		///fmt.Println("ok", allExtensions)

		// Store the path and extension in the map
		for i := 0; i < len(allExtensions); i++ {
			if allExtensions[0] != allExtensions[i] {
				data[line] = allExtensions[i]
			}

		}

	}

	// Check for any errors that occurred during the scan
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return data
}
