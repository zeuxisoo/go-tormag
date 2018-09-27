package utils

import (
    "os"
)

// IsDirectory to check the path is or not directory
func IsDirectory(path string) bool {
    fileInfo, err := os.Stat(path)
	if err != nil {
		return false
    }

	return fileInfo.IsDir() == true
}

// CreateOrIsDirectoryExists to create directory if the path is not directory
func CreateOrIsDirectoryExists(path string) bool {
    // If the path is not directory, create it
    if IsDirectory(path) == false {
        err := os.MkdirAll(path, os.ModePerm)
        return err == nil
    }

    return true
}
