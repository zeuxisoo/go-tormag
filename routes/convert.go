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

    if utils.CreateOrIsDirectoryExists(setting.AttachmentPath) == false {
        message = "Cannot create storage directory"
    }else{
        file, err := c.FormFile("file")

        if err != nil {
            message = "Cannot read the file from form data"
        }else if utils.IsTorrentFile(file.Filename) == false {
            message = "The upload file is not torrent file"
        }else if utils.IsBiggerFile(file.Size, setting.AttachmentMaxSize) == true {
            message = fmt.Sprintf("The upload file is bigger than %dMB", setting.AttachmentMaxSize)
        }else{
            // TODO: move and convert file

            ok      = true
            message = "Successfully, File converted"
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "ok"     : ok,
        "message": message,
    })
}
