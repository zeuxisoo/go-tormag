package routes

import (
    "fmt"
	"net/http"

    "github.com/gin-gonic/gin"

    "github.com/zeuxisoo/go-tormag/pkg/setting"
    "github.com/zeuxisoo/go-tormag/pkg/utils"
)

// ConvertPost return the convert state
func ConvertPost(c *gin.Context) {
    ok      := false
    message := ""

    file, err := c.FormFile("file")
    if err != nil {
        message = "Cannot read the file from form data"
    }else{
        fmt.Println(file.Filename)

        if utils.IsTorrentFile(file.Filename) == false {
            message = "The upload file is not torrent file"
        }else if utils.IsBiggerFile(file.Size, setting.AttachmentMaxSize) == true {
            message = fmt.Sprintf("The upload file is bigger than %dMB", setting.AttachmentMaxSize)
        }else if utils.CreateOrIsDirectoryExists(setting.AttachmentPath) == false {
            message = "Cannot create storage directory"
        }else{
            ok      = true
            message = "Successfully, File converted"
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "ok"     : ok,
        "message": message,
    })
}
