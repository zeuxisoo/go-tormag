package cmd

import (
    "github.com/urfave/cli"

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
        stringFlag("output, o", "bigger.txt", "Output file to store the bigger file names"),
    },
}

func runBigger(c *cli.Context) error {
    logger.Infof("TODO: complete the bigger function")

    return nil
}
