package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

// Web command
var Web = cli.Command{
	Name:        "web",
	Usage:       "Start web interface",
	Description: "Run web server to start the web application",
	Action:      runWeb,
	Flags: []cli.Flag{
		stringFlag("address, a", "127.0.0.1", "Server address"),
		stringFlag("port, p", "3000", "Server port"),
	},
}

func runWeb(c *cli.Context) error {
	fmt.Println("TODO: run web server")
	return nil
}
