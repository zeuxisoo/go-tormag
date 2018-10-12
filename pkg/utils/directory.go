package utils

import (
    "os"
)

// IsDirectoryExists to check the path is or not exists
func IsDirectoryExists(path string) bool {
    _, err := os.Stat(path)

    // Must IsNotExist
    // - https://github.com/golang/go/blob/master/src/os/error.go#L76
    // - https://github.com/golang/go/blob/master/src/os/error_unix.go#L11-L19
    return os.IsNotExist(err) == false
}

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
