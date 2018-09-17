package render

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gin-gonic/gin/render"
	"gopkg.in/flosch/pongo2.v3"

	"github.com/zeuxisoo/go-tormag/pkg/view"
)

// Pongo2Render object
type Pongo2Render struct {
	name	string
	data	gin.H
}

// NewPongo2Render return the instance
func NewPongo2Render() *Pongo2Render {
	return &Pongo2Render{}
}

// Instance return a new Pongo2Render struct per request
func (p Pongo2Render) Instance(name string, data interface{}) render.Render {
	return Pongo2Render{
		name: name,
		data: data.(gin.H),
	}
}

/*
 * Implement for the render.Render
 */

// Render the template to the response
func (p Pongo2Render) Render(w http.ResponseWriter) error {
	assetBytes, err := view.Asset(p.name)
	if err != nil {
		return err
	}

	tpl, err := pongo2.FromString(string(assetBytes))
	if err != nil {
		return err
	}

	p.WriteContentType(w)

	return tpl.ExecuteWriter(pongo2.Context(p.data), w)
}

// WriteContentType add the content type header to the response
func (p Pongo2Render) WriteContentType(w http.ResponseWriter) {
	header := w.Header()

	if contentType := header["Content-Type"]; len(contentType) == 0 {
		header["Content-Type"] = []string{ "text/html; charset=utf-8" }
	}
}
