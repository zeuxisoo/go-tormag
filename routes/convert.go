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
        }else if md5, err := utils.GetFileMd5(file); err != nil {
            message = fmt.Sprintf("Cannot get the md5 hash from upload file: %v", err)
        }else{
            fmt.Println(md5)

            ok      = true
            message = "Successfully, File converted"
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "ok"     : ok,
        "message": message,
    })
}
