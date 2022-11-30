package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	cp "github.com/otiai10/copy"
)

func main() {
	// Get a list of all folders in the current directory
	folders, err := GetFolders(".")
	if err != nil {
		panic(err)
	}

	_, largestNumber := GetLatestFolder(folders)

	// Make a copy of the folder `Folder_template` and name it `Invoice {largestNumber + 1}`
	newName := fmt.Sprintf("Invoice %d", largestNumber+1)
	err = cp.Copy("Folder_template", newName)
	if err != nil {
		panic(err)
	}

	// Rename the file in the new folder from `Invoice_template.xlsx` to `Invoice {largestNumber + 1}.xlsx`
	err = os.Rename(fmt.Sprintf("%s/Invoice_template.xlsx", newName), fmt.Sprintf("%s/Invoice %d.xlsx", newName, largestNumber+1))
	if err != nil {
		panic(err)
	}
}

// GetFolders returns a list of all folders in the given directory
func GetFolders(dir string) ([]string, error) {
	var folders []string

	// Read the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Loop through the files
	for _, file := range files {
		// If the file is a directory, add it to the list
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}

	return folders, nil
}

func GetLatestFolder(folders []string) (string, int) {
	// Get the folder ending with the highest number
	// Folders are named like `Invoice 1`, `Invoice 23`, etc.
	var latestFolder string
	var latestFolderNumber int
	for _, folder := range folders {
		// Split the folder name into two parts at the final space
		parts := SplitLast(folder, " ")
		if len(parts) != 2 {
			continue
		}
		// Convert the second part to an integer
		number, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		// If the number is higher than the latest number, set it as the latest
		if number > latestFolderNumber {
			latestFolder = folder
			latestFolderNumber = number
		}
	}

	return latestFolder, latestFolderNumber
}

func SplitLast(s, sep string) []string {
	parts := []string{s}
	if i := strings.LastIndex(s, sep); i >= 0 {
		parts = []string{s[:i], s[i+len(sep):]}
	}
	return parts
}
