package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeGet return the index page
func HomeGet(c *gin.Context) {
	c.HTML(http.StatusOK, "views/home/index.html", gin.H{
		"sumLabel": "Sum value",
	})
}

// AboutGet return the about page
func AboutGet(c *gin.Context) {
	c.HTML(http.StatusOK, "views/home/about.html", gin.H{})
}
