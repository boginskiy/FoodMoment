package logg

import (
	"log"
	"os"
)

func CreateFolder(nameFolder string) error {
	if _, err := os.Stat(nameFolder); os.IsNotExist(err) {
		// Creating the logs folder, if it does not exist
		return os.MkdirAll(nameFolder, 0755)
	}
	return nil
}

func CreateFile(nameFile string) (*os.File, error) {
	return os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func MakeDirAndFile(nameFolder, path string) *os.File {
	// Create folder
	err := CreateFolder(nameFolder)
	if err != nil {
		log.Fatal(err)
	}

	// Create files
	file, err := CreateFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
