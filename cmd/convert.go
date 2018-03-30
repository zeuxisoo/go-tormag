package cmd

import (
    "github.com/urfave/cli"

    "github.com/zeuxisoo/go-tormag/util/logger"
)

var Convert = cli.Command{
    Name:  "convert",
    Usage: "Convert the torrent to magnet",
    Description: "Enter the torrent directory, then convert the torrent files to magnet link from the input directory",
    Action: runConvert,
    Flags: []cli.Flag{
        stringFlag("directory, d", "", "Torrent directory"),
    },
}

func runConvert(c *cli.Context) error {
    if !c.IsSet("directory") {
        printCommandHelpAndExit(c, "convert")
    }

    logger.Infof("Directory: %s", c.String("directory"))

    return nil
}
