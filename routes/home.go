package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeGet return the index page
func HomeGet(c *gin.Context) {
	c.HTML(http.StatusOK, "views/home/index.html", gin.H{})
}
