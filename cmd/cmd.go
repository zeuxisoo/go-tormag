package cmd

import (
    "github.com/urfave/cli"
)

func stringFlag(name, value, usage string) cli.StringFlag {
    return cli.StringFlag{
        Name :  name,
        Value: value,
        Usage: usage,
    }
}

func printCommandHelpAndExit(c *cli.Context, command string) {
    cli.ShowCommandHelpAndExit(c, command, 0)
}
