package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	ginStatic "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"gopkg.in/flosch/pongo2.v3"

	"github.com/zeuxisoo/go-tormag/internal/render"
	"github.com/zeuxisoo/go-tormag/internal/setting"
	"github.com/zeuxisoo/go-tormag/public"
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

    engine.Use(func(c *gin.Context) {
        // Define service worker scope to support /static/sw.js
        if c.Request.URL.Path == "/static/sw.js" {
            c.Header("Service-Worker-Allowed", "/")
        }
    })

    engine.Use(ginStatic.Serve("/static", public.NewFileSystem("static")))

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
    if strings.ToLower(setting.Mode) != "dev" {
        fmt.Printf("Running on %s mode\n", setting.Mode)
        fmt.Printf("Listening and serving HTTP on %s\n", appURL)
    }

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

    //
    engine.GET("/robots.txt", func(c *gin.Context) {
        data, _ := public.Files.ReadFile("static/robots.txt")

        c.String(200, string(data))
    })

    //
    engine.POST("/convert", routes.ConvertPost)
    engine.POST("/bigger", routes.BiggerPost)
    engine.POST("/about", routes.AboutPost)

    // Handle 404 not found problem in frontend when route change and then refresh page by pass all request to index view/template
    engine.NoRoute(routes.HomeGet)
}
