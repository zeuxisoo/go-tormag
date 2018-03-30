package cmd

import (
    "os"
    "fmt"
    "io/ioutil"
    "path/filepath"
    "strings"

    "github.com/urfave/cli"
    "github.com/anacrolix/torrent/metainfo"

    "github.com/zeuxisoo/go-tormag/util/logger"
)

var Convert = cli.Command{
    Name:  "convert",
    Usage: "Convert the torrent to magnet",
    Description: "Enter the torrent directory, then convert the torrent files to magnet link from the input directory",
    Action: runConvert,
    Flags: []cli.Flag{
        stringFlag("directory, d", "", "Torrent directory"),
        stringFlag("output, o", "magnets.txt", "Output file to store the converted magnet links"),
    },
}

func runConvert(c *cli.Context) error {
    if !c.IsSet("directory") {
        printCommandHelpAndExit(c, "convert")
    }

    directory   := c.String("directory")
    output_file := c.String("output")

    logger.Infof("Directory path : %s", directory)

    //
    files, err := ioutil.ReadDir(directory)
    if err != nil {
        logger.Fatalf("Error, Cannot read the files in torrent folder path (%s)", err)
    }

    //
    var torrent_path string
    var magnets []string

    for _, f := range files {
        if (f.IsDir() == false && filepath.Ext(f.Name()) == ".torrent") {
            torrent_path = filepath.Join(directory, f.Name())

            mi, err := metainfo.LoadFromFile(torrent_path)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error, Cannot read the metainfo from file (%s)\n", err)
                continue
            }

            info, err := mi.UnmarshalInfo()
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error, Cannot unmarshalling the metainfo from file (%s)\n", err)
                continue
            }

            magnets = append(magnets, mi.Magnet(info.Name, mi.HashInfoBytes()).String())
        }
    }

    ioutil.WriteFile(output_file, []byte(strings.Join(magnets, "\n")), 0644)

    //
    logger.Infof("Output filename: %s", output_file)

    return nil
}
