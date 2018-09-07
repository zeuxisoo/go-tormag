package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/zeuxisoo/go-tormag/cmd"
)

const APP_VER = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "Tormag"
	app.Usage = "A torrent to magnet and rename tools"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.Convert,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
