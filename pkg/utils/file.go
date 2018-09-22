package utils

import (
    "os"
)

// IsFile return the file path is or not file
func IsFile(filePath string) bool {
    fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
    }

	return fileInfo.IsDir() == false
}
