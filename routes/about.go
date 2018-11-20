package routes

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/zeuxisoo/go-tormag/pkg/setting"
)

// AboutPost return the application infomation
func AboutPost(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "name": "Tormag",
        "mode": setting.Mode,
        "build": map[string]string{
            "time": setting.BuildTime,
            "hash": setting.BuildHash,
        },
    })
}
