package projectfileslocation

// Projectfileslocation contains the principal engine to search the files base on the extension

import (
	"bufio"
	"fmt"
	"log"

	"os"
	"path/filepath"

	"github.com/kakaw2016/Gotraining/Ariel/configurationFile"
)

// Files located using the extensions

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

// Magor file path search in Directories

func DirectoryFilesLocation() {

	//Read input from the configuration file

	currentDir, _ := os.Getwd()
	configPath := filepath.Join(currentDir, "configurationfile.txt")
	relativePath, _ := filepath.Rel(currentDir, configPath)
	fmt.Println("Reading config file at relative path:", relativePath)

	configMap := configurationFile.ReadConfigurationFile(relativePath)

	//Save search file path in the output files and update the line each time the research is launched.

	externalStoredata, err := os.OpenFile("FileSearchFinalResult.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer externalStoredata.Close()

	w := bufio.NewWriter(externalStoredata)

	for filePaths, fileExtension := range configMap {

		storing0 := fmt.Sprintf("\n--\nLocal path to the files: %s, the extension selected: %s\n", filePaths, fileExtension)
		_, _ = w.WriteString(storing0)

		extension := "*" + fileExtension

		//Complete the repertory search with the key factors filepaths and extension

		files, err := filesFilteredbyExtension(filePaths, extension)
		if err != nil {

			storing1 := fmt.Sprint("filesFilteredbyExtension has an Error", err, "\n")
			_, _ = w.WriteString(storing1)
		} else if files == nil {

			storing2 := fmt.Sprint("\n", "No File path can be found after search.", "\n")
			_, _ = w.WriteString(storing2)

		} else {
			for _, locatePathFiles := range files {

				storing3 := fmt.Sprint("**\n", locatePathFiles, "\n")
				_, _ = w.WriteString(storing3)

			}

		}

	}
	fmt.Println("Please Check Your Final Results Inside File Search Text File.")

	fmt.Println("Program Ends")
	w.Flush()
}
