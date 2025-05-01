package utils

import (
	"os"
)

func IsValidFile(file_path string) bool {
	// Check if the file exists.
	info, err := os.Stat(file_path)
	if err != nil {
		return false // File does not exist/not accessible, hence false.
	}
	return !info.IsDir() // Check if the file is not a directory.
}
