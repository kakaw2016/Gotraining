package configurationFile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Configuration details in file that generates all the input we will use to find the files

func ReadConfigurationFile(configurationFilepath string) (map[string]string, error) {

	// Open the config file in read-only mode
	configFile, err := os.Open(configurationFilepath)
	if err != nil {
		fmt.Println(err)

	}
	defer configFile.Close()

	// Create scanner for config.txt
	scanner := bufio.NewScanner(configFile)

	// Create mapping to store paths
	valueStorage := make(map[string]string)
	dataStorage := make(map[string]string)

	// Create slice to store extensions

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if key == "path" {
				// Add path to mapping
				valueStorage["path"] = value

			} else if key == "extensions" {

				valueStorage["extensions"] = value
			}
		}
		dataStorage[valueStorage["path"]] = valueStorage["extensions"]

	}

	// Check for errors while scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dataStorage, nil
}
