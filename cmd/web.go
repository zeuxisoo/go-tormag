package cmd

import (
    "fmt"
    "time"
    "strings"

	"github.com/urfave/cli"
    "github.com/gin-gonic/gin"
    ginStatic "github.com/gin-contrib/static"
    "github.com/gin-contrib/cors"
	"gopkg.in/flosch/pongo2.v3"

    "github.com/zeuxisoo/go-tormag/pkg/setting"
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
        stringFlag("config, c", "", "Custom server config file"),
	},
}

func runWeb(c *cli.Context) {
    // Set custom config file
    if c.IsSet("config") {
        setting.CustomConfig = c.String("config")
    }

    // Init settings from config file
    setting.NewSetting()

    // Overwrite the custom argument to settings
    if c.IsSet("address") {
        setting.Address = c.String("address")
    }

    if c.IsSet("port") {
        setting.Port = c.String("port")
    }

    if c.IsSet("mode") {
        setting.Mode = c.String("mode")
    }

	//
	appURL := fmt.Sprintf("%s:%s", setting.Address, setting.Port)

	//
	if strings.ToLower(setting.Mode) == "dev" {
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}

	//
    engine := gin.Default()
    engine.MaxMultipartMemory = setting.AttachmentMaxSize << 20 // Set a lower memory limit for multipart (default: 4MB, app.ini: 8MB)

    engine.Use(ginStatic.Serve("/static", static.NewFileSystem("static")))

    if setting.CrossOriginEnable == true {
        engine.Use(cors.New(cors.Config{
            AllowOrigins    : setting.CrossOriginAllowOrigins,
            AllowMethods    : setting.CrossOriginAllowMethods,
            AllowHeaders    : setting.CrossOriginAllowHeaders,
            ExposeHeaders   : setting.CrossOriginExposeHeaders,
            AllowCredentials: setting.CrossOriginAllowCredentials,
            AllowOriginFunc : func(origin string) bool {
                return true
            },
            MaxAge          : time.Duration(setting.CrossOriginMaxAge) * time.Hour,
        }))
    }

    engine.Use(gin.Recovery())

    //
	registerRender(engine)
	registerRoutes(engine)

    //
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
