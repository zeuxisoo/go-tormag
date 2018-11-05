package cmd

import (
    "fmt"
    "strings"
    "io/ioutil"
    "path/filepath"

    "github.com/urfave/cli"
    "github.com/anacrolix/torrent/metainfo"

    "github.com/zeuxisoo/go-tormag/pkg/logger"
)

// Bigger command
var Bigger = cli.Command{
    Name:        "bigger",
    Usage:       "Find bigger file in the torrent file",
    Description: "Show the bigger file name from the input torrent file",
    Action:      runBigger,
    Flags: []cli.Flag{
        stringFlag("directory, d", "", "A torrent directory"),
        stringFlag("file, f", "", "A single torrent file"),
        stringFlag("output, o", "biggers.txt", "Output file to store the bigger file names"),
    },
}

func runBigger(c *cli.Context) error {
    if !c.IsSet("directory") && !c.IsSet("file") {
        printCommandHelpAndExit(c, "bigger")
    }

    directory  := c.String("directory")
    file       := c.String("file")
    outputFile := c.String("output")

    var result []byte

    if c.IsSet("directory") {
        result = findBiggerFilesFromDirectory(directory)
    } else if c.IsSet("file") {
        result = findBiggerFileFromFile(file)
    } else {
        result = []byte("")
    }

    if err := ioutil.WriteFile(outputFile, result, 0644); err != nil {
        logger.Fatalf("Error: Cannot write the result to the output file")
    }

    //
    logger.Infof("Output: %s", outputFile)

    return nil
}

func findBiggerFilesFromDirectory(directory string) []byte {
    //
    logger.Infof("Directory: %s", directory)

    files, err := ioutil.ReadDir(directory)
    if err != nil {
        logger.Fatalf("Error: Cannot read the files in torrent folder path (%s)\n", err)
    }

    //
    logger.Infof("=> Reading ... ...")

    fullPath     := ""
    torrentInfos := make(map[string]metainfo.Info)

    for _, file := range files {
        if (file.IsDir() == false && filepath.Ext(file.Name()) == ".torrent") {
            fullPath = filepath.Join(directory, file.Name())

            mi, err := metainfo.LoadFromFile(fullPath)
            if err != nil {
                printBiggerTorrentError("Cannot read the metainfo from file", fullPath, err)
                continue
            }

            info, err := mi.UnmarshalInfo()
            if err != nil {
                printBiggerTorrentError("Cannot unmarshalling the metainfo from file", fullPath, err)
                continue
            }

            torrentInfos[fullPath] = info
        }
    }

    //
    logger.Infof("=> Finding ... ...")

    var biggerFiles []string

    for fullPath, info := range torrentInfos {
        biggerFiles = append(
            biggerFiles,
            fmt.Sprintf("%s => %s", fullPath, findBiggerFileName(info)),
        )
    }

    return []byte(strings.Join(biggerFiles, "\n"))
}

func findBiggerFileFromFile(file string) []byte {
    logger.Infof("File: %s", file)

    mi, err := metainfo.LoadFromFile(file)
    if err != nil {
        printBiggerTorrentError("Cannot read the metainfo from file", file, err)
    }

    info, err := mi.UnmarshalInfo()
    if err != nil {
        printBiggerTorrentError("Cannot unmarshal the metainfo from file", file, err)
    }

    biggerFile := fmt.Sprintf("%s => %s", file, findBiggerFileName(info))

    return []byte(biggerFile)
}

//
func findBiggerFileName(info metainfo.Info) string {
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
