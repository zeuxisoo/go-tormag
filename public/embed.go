package public

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed **/*
var Files embed.FS

// FileSystem object
type FileSystem struct {
    fileSystem http.FileSystem
}

// Open the file by name
func (p *FileSystem) Open(name string) (http.File, error) {
    return p.fileSystem.Open(name)
}

// Exists to check the file is or exists
func (p *FileSystem) Exists(prefix string, filePath string) bool {
    if url := strings.TrimPrefix(filePath, prefix); len(url) < len(filePath) {
        if _, err := p.fileSystem.Open(url); err != nil {
            return false
        }

        return true
    }

    return false
}

// NewFileSystem return FileSystem instance object
func NewFileSystem(root string) *FileSystem {
    fileSystem, err := fs.Sub(Files, root)
    if err != nil {
        panic(err)
    }

    return &FileSystem{
        fileSystem: http.FS(fileSystem),
    }
}
