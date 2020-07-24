package utils

import (
    "github.com/anacrolix/torrent/metainfo"
)

// FindBigFilenameFromTorrentMetaInfo return the bigger file name from meta info object
func FindBigFilenameFromTorrentMetaInfo(info metainfo.Info) string {
    var currentLength int64
    var currentFile metainfo.FileInfo

    for _, file := range info.Files {
        if file.Length > currentLength || currentLength == 0 {
            currentLength = file.Length
            currentFile   = file
        }
    }

    name := info.Name
    if len(currentFile.Path) >= 1 {
        name = currentFile.Path[0]
    }

    return name
}
