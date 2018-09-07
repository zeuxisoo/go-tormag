package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/urfave/cli"

	"github.com/zeuxisoo/go-tormag/util/logger"
)

var Convert = cli.Command{
	Name:        "convert",
	Usage:       "Convert the torrent to magnet",
	Description: "Enter the torrent directory, then convert the torrent files to magnet link from the input directory",
	Action:      runConvert,
	Flags: []cli.Flag{
		stringFlag("directory, d", "", "A torrent directory"),
		stringFlag("file, f", "", "A single torrent file"),
		stringFlag("output, o", "magnets.txt", "Output file to store the converted magnet links"),
	},
}

func runConvert(c *cli.Context) error {
	if !c.IsSet("directory") && !c.IsSet("file") {
		printCommandHelpAndExit(c, "convert")
	}

	directory := c.String("directory")
	file := c.String("file")
	output_file := c.String("output")

	var result []byte

	if c.IsSet("directory") {
		result = convertDirectoryToMagentBytes(directory)
	} else if c.IsSet("file") {
		result = convertFileToMagnetBytes(file)
	} else {
		result = []byte("")
	}

	if err := ioutil.WriteFile(output_file, result, 0644); err != nil {
		logger.Fatalf("[Error] >> Cannot write the result to the output file")
	}

	//
	logger.Infof("[Output] >> %s", output_file)

	return nil
}

func convertDirectoryToMagentBytes(directory string) []byte {
	logger.Infof("[Directory] >> %s", directory)

	//
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		logger.Fatalf("[Error] >> Cannot read the files in torrent folder path (%s)", err)
	}

	//
	var torrent_path string
	var magnets []string

	for _, f := range files {
		if f.IsDir() == false && filepath.Ext(f.Name()) == ".torrent" {
			torrent_path = filepath.Join(directory, f.Name())

			mi, err := metainfo.LoadFromFile(torrent_path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[Error] >> Cannot read the metainfo from file (%s)\n", err)
				continue
			}

			info, err := mi.UnmarshalInfo()
			if err != nil {
				fmt.Fprintf(os.Stderr, "[Error] >> Cannot unmarshalling the metainfo from file (%s)\n", err)
				continue
			}

			magnets = append(magnets, mi.Magnet(info.Name, mi.HashInfoBytes()).String())
		}
	}

	return []byte(strings.Join(magnets, "\n"))
}

func convertFileToMagnetBytes(file string) []byte {
	logger.Infof("[File] >> %s", file)

	mi, err := metainfo.LoadFromFile(file)
	if err != nil {
		logger.Fatalf("[Error] >> Cannot read the metainfo from file (%s)\n", err)
	}

	info, err := mi.UnmarshalInfo()
	if err != nil {
		logger.Fatalf("[Error] >> Cannot unmarshalling the metainfo from file (%s)\n", err)
	}

	magnet := mi.Magnet(info.Name, mi.HashInfoBytes()).String()

	return []byte(magnet)
}
