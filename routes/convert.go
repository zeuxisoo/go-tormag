package routes

import (
    "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

        // TODO: store the file

        ok      = true
        message = "Successfully, File converted"
    }

    c.JSON(http.StatusOK, gin.H{
        "ok"     : ok,
        "message": message,
    })
}
