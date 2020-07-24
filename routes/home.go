package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-tormag/internal/setting"
	"github.com/zeuxisoo/go-tormag/internal/utils"
)

// HomeGet return the index page
func HomeGet(c *gin.Context) {
    homeIndexViewFile := "views/home/index.html"

    if setting.BuildEnv != "development" || utils.IsFileExists(homeIndexViewFile) == true {
        c.HTML(http.StatusOK, homeIndexViewFile, gin.H{})
    }else{
        c.String(http.StatusOK, "Hello world in develop mode")
    }
}
