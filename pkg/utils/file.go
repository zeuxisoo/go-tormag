package utils

import (
    "os"
    "fmt"
    "path"
    "strings"
    "mime/multipart"
    "io/ioutil"
    "crypto/md5"
    "encoding/hex"
)

// GetFileExtension return the file extension
func GetFileExtension(filePath string) string {
    return path.Ext(filePath)
}

// GetFileMD5 return a MD5 hashed string
func GetFileMD5(fileHeader *multipart.FileHeader) (string, error) {
    file, err := fileHeader.Open()
    defer file.Close()
    if err != nil {
        return "", err
    }

    contentBytes, err := ioutil.ReadAll(file)
    if err != nil {
        return "", err
    }

    hash := md5.New()
    hash.Write(contentBytes)

    return hex.EncodeToString(hash.Sum(nil)), nil
}

// IsFile return the file path is or not file
func IsFile(filePath string) bool {
    fileInfo, err := os.Stat(filePath)
    if err != nil {
        return false
    }

    return fileInfo.IsDir() == false
}

// IsMatchFileExtension return the source extension is or not match target extension
func IsMatchFileExtension(sourceExtension string, targetExtension string) bool {
    return strings.ToLower(sourceExtension) == strings.ToLower(targetExtension)
}

// IsTorrentFile return the filePath is or not torrent file
func IsTorrentFile(filePath string) bool {
    extension := GetFileExtension(filePath)

    return IsMatchFileExtension(".torrent", extension)
}

// IsBiggerFile return the file size (Bytes) is or not bigger than setting.MAX_SIZE (MegaBytes) size
func IsBiggerFile(fileSize int64, defaultFileSize int64) bool {
    // Convert Bytes to KB
    fileSizeInKB := fileSize / 1024

    // Convert MB to KB
    defaultSizeInKB := defaultFileSize * 1024

    return fileSizeInKB > defaultSizeInKB
}

// RemoveFile to make sure the given path is file before remove it
func RemoveFile(path string) error {
    if IsFile(path) == true {
        return os.Remove(path)
    }

    return fmt.Errorf("Path is not file: %s", path)
}
