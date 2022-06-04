package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadFile reads the file and returns the content
func ReadFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data
}

// WriteFile writes the content to the file
func WriteFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		panic(err)
	}

	return nil
}

// Auto join the path
// If the path is not absolute, it will be joined with the current working directory
func JoinPaths(paths ...string) string {
	return filepath.Join(paths...)
}
