package views

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gin-gonic/gin/render"
	"github.com/gobuffalo/plush"
)

// PlushRender object
type PlushRender struct {
	name 	string
	context *plush.Context
}

// NewPlushRender return the instance
func NewPlushRender() *PlushRender {
	return &PlushRender{}
}

// Instance return a new PlushRender struct per request
func (p PlushRender) Instance(name string, data interface{}) render.Render {
	return PlushRender{
		name   : name,
		context: plush.NewContextWith(data.(gin.H)),
	}
}

/*
 * Implement for the render.Render
 */

// Render the template to the response
func (p PlushRender) Render(w http.ResponseWriter) error {
	assetBytes, err := Asset(string(p.name))
	if err != nil {
		return err
	}

	result, err := plush.Render(string(assetBytes), p.context)
	if err != nil {
		return err
	}

	p.WriteContentType(w)

	w.Write([]byte(result))

	return nil
}

// WriteContentType add the content type header to the response
func (p PlushRender) WriteContentType(w http.ResponseWriter) {
	header := w.Header()

	if contentType := header["Content-Type"]; len(contentType) == 0 {
		header["Content-Type"] = []string{ "text/html; charset=utf-8" }
	}
}
