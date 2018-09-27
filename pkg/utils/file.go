package utils

import (
	"strings"
    "os"
    "path"
)

// GetFileExtension return the file extension
func GetFileExtension(filePath string) string {
    return path.Ext(filePath)
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
