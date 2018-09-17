package cmd

import (
	"github.com/urfave/cli"

	"github.com/zeuxisoo/go-tormag/pkg/logger"
)

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func printCommandHelpAndExit(c *cli.Context, command string) {
	cli.ShowCommandHelpAndExit(c, command, 0)
}


func printConvertTorrentError(message string, torrentPath string, err error) {
	logger.Errorf("[Error]")
	logger.Errorf("=> %s", message)
	logger.Errorf("=> file : %s", torrentPath)
	logger.Errorf("=> error: %s", err)
}
