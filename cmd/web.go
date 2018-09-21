package cmd

import (
    "fmt"
    "time"

	"github.com/urfave/cli"
    "github.com/gin-gonic/gin"
    ginStatic "github.com/gin-contrib/static"
    "github.com/gin-contrib/cors"
	"gopkg.in/flosch/pongo2.v3"

    "github.com/zeuxisoo/go-tormag/pkg/render"
    "github.com/zeuxisoo/go-tormag/pkg/static"
	"github.com/zeuxisoo/go-tormag/routes"
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
    engine.MaxMultipartMemory = 8 << 20 // Set a lower memory limit for multipart forms to 8MB (default is 32MB)

    engine.Use(ginStatic.Serve("/static", static.NewFileSystem("static")))
    engine.Use(cors.New(cors.Config{
		AllowOrigins    : []string{"http://127.0.0.1:8080"},    // TODO: move to config
		AllowMethods    : []string{"POST"},
		AllowHeaders    : []string{"Origin"},
		ExposeHeaders   : []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc : func(origin string) bool {
			return true
		},
		MaxAge          : 12 * time.Hour,
	}))
    engine.Use(gin.Recovery())

	registerRender(engine)
	registerRoutes(engine)

	engine.Run(appURL)
}

func registerRender(engine *gin.Engine) {
	// engine.HTMLRender = render.NewBaseRender(&render.BaseOption{
	// 	Functions: render.BaseFunctions{
	// 		// e.g. {{ says "Hello world" ":D" }} ==> Hello world - Check - :D
	// 		"says": func(value ...interface{}) string {
	// 			helloWorld  := value[0].(string)
	// 			suffixWorld := value[1].(string)

	// 			return helloWorld + " - Check - " + suffixWorld
	// 		},
	// 	},
	// })

	// engine.HTMLRender = render.NewPlushRender(&render.PlushOption{
	// 	Functions: render.PlushFunctions{
	//		// e.g. <%= says("Hello world", ":D") %> ==> Hello world - Check - :D
	// 		"says": func(value string, argument string) string {
	// 			return value + " - Check - " + argument
	// 		},
	// 	},
	// })

	engine.HTMLRender = render.NewPongo2Render(&render.Pongo2Option{
		Filters: render.Pongo2FilterFunctions{
			"says": func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
				value    := in.String()
				argument := param.String()

				return pongo2.AsValue(value + " - Check - " + argument), nil
			},
		},
	})
}

func registerRoutes(engine *gin.Engine) {
    engine.GET("/", routes.HomeGet)
    engine.POST("/convert", routes.ConvertPost)
}
