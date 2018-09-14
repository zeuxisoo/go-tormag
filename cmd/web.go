package cmd

import (
	"fmt"
	"html/template"

	"github.com/urfave/cli"
	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-tormag/web/views"
	"github.com/zeuxisoo/go-tormag/web/controllers"
)

// Web command
var Web = cli.Command{
	Name:        "web",
	Usage:       "Start web interface",
	Description: "Run web server to start the web application",
	Action:      runWeb,
	Flags: []cli.Flag{
		stringFlag("address, a", "0.0.0.0", "Server address"),
		stringFlag("port, p", "3000", "Server port"),
		stringFlag("mode, m", "dev", "Server mode"),
	},
}

func runWeb(c *cli.Context) {
	//
	appURL := fmt.Sprintf("%s:%s", c.String("address"), c.String("port"))

	//
	if c.String("mode") == "dev" {
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}

	//
	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.SetFuncMap(template.FuncMap{})

	registerRender(engine)
	registerRoutes(engine)

	engine.Run(appURL)
}

func registerRender(engine *gin.Engine) {
	engine.HTMLRender = views.NewRender()
}

func registerRoutes(engine *gin.Engine) {
	engine.GET("/", controllers.HomeGet)
	engine.GET("/about", controllers.AboutGet)
}
