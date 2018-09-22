package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/zeuxisoo/go-tormag/cmd"
)

//go:generate go-bindata -pkg view -o pkg/view/view.go -ignore=.go views/...
//go:generate go-bindata -pkg static -o pkg/static/static.go -ignore=.go static/...
//go:generate go-bindata -pkg config -o config/config.go -ignore=.go  config/...

// AppVersion number
const AppVersion = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "Tormag"
	app.Usage = "A torrent to magnet and rename tools"
	app.Version = AppVersion
	app.Commands = []cli.Command{
		cmd.Convert,
		cmd.Web,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
