package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/zeuxisoo/go-tormag/cmd"
)

// AppVersion number
const AppVersion = "0.2.0"

func main() {
	app := cli.NewApp()
	app.Name = "Tormag"
	app.Usage = "A torrent to magnet and rename tools"
	app.Version = AppVersion
	app.Commands = []cli.Command{
        cmd.Bigger,
        cmd.Convert,
		cmd.Web,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
