package routes

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-tormag/internal/setting"
	"github.com/zeuxisoo/go-tormag/internal/utils"
)

// BiggerPost return the convert state
func BiggerPost(c *gin.Context) {
    ok      := false
    message := ""
    data    := map[string]string{
        "id"  : fmt.Sprintf("%s-%d-%s", utils.DateFormat(time.Now(), "YYYYMMDDHHmmss"), utils.Timestamp(), utils.RandomString(8)),
        "file": "",
        "md5" : "",
        "big" : "",
    }

    if !utils.CreateOrIsDirectoryExists(setting.AttachmentPath) {
        message = "Cannot create storage directory"
    }else{
        file, err := c.FormFile("file")
        if err == nil {
            // Set filename to file field on each response message when upload without parse error
            // - like multipart/form-data field name incorrect (e.g. "file[]" not "file")
            data["file"] = file.Filename
        }

        if err != nil {
            message = "Cannot read the file from form data"
        }else if !utils.IsTorrentFile(file.Filename) {
            message = "The upload file is not torrent file"
        }else if utils.IsBiggerFile(file.Size, setting.AttachmentMaxSize) {
            message = fmt.Sprintf("The upload file is bigger than %dMB", setting.AttachmentMaxSize)
        }else if fileMD5, err := utils.GetFileMD5(file); err != nil {
            message = fmt.Sprintf("Cannot get the md5 hash from upload file: %v", err)
        }else{
            fileExtension := utils.GetFileExtension(file.Filename)
            fileName      := fileMD5 + fileExtension
            fileFullPath  := path.Join(setting.AttachmentPath, fileName)

            if err := c.SaveUploadedFile(file, fileFullPath); err != nil {
                message = "Cannot save uploaded file"
            }else if mi, err := metainfo.LoadFromFile(fileFullPath); err != nil {
                message = "Cannot read the metainfo from uploaded file"
            }else if  info, err := mi.UnmarshalInfo(); err != nil {
                message = "Cannot unmarshal the metainfo from uploaded file"
            }else{
                ok      = true
                message = "Successfully, File found"

                data["md5"] = fileMD5
                data["big"] = utils.FindBigFilenameFromTorrentMetaInfo(info)

                utils.RemoveFile(fileFullPath)
            }
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "ok"     : ok,
        "message": message,
        "data"   : data,
    })
}
