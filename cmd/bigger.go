package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/urfave/cli"

	"github.com/zeuxisoo/go-tormag/internal/logger"
	"github.com/zeuxisoo/go-tormag/internal/utils"
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

    var result string

    if c.IsSet("directory") {
        result = findBiggerFilesFromDirectory(directory)
    }else if c.IsSet("file") {
        info, err := findBiggerFileFromFile(file)
        if err != nil {
            logger.Fatalf("Error: %s", err)
        }

        result = fmt.Sprintf("%s => %s", file, utils.FindBigFilenameFromTorrentMetaInfo(info))
    }else{
        result = ""
    }

    if err := ioutil.WriteFile(outputFile, []byte(result), 0644); err != nil {
        logger.Fatalf("Error: Cannot write the result to the output file")
    }

    //
    logger.Infof("Output: %s", outputFile)

    return nil
}

func findBiggerFilesFromDirectory(directory string) string {
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

            info, err := findBiggerFileFromFile(fullPath)
            if err != nil {
                logger.Errorf("Error: %s", err)
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
            fmt.Sprintf("%s => %s", fullPath, utils.FindBigFilenameFromTorrentMetaInfo(info)),
        )
    }

    return strings.Join(biggerFiles, "\n")
}

func findBiggerFileFromFile(file string) (metainfo.Info, error) {
    logger.Infof("File: %s", file)

    defaultMetaInfo := metainfo.Info{}

    mi, err := metainfo.LoadFromFile(file)
    if err != nil {
        return defaultMetaInfo, fmt.Errorf("Cannot read the metainfo from file: %s. %v", file, err)
    }

    info, err := mi.UnmarshalInfo()
    if err != nil {
        return defaultMetaInfo, fmt.Errorf("Cannot unmarshal the metainfo from file: %s. %v", file, err)
    }

    return info, nil
}
