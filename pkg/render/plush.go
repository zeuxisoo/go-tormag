package render

import (
	"net/http"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/gobuffalo/plush"

	"github.com/zeuxisoo/go-tormag/pkg/view"
)

// PlushFunctions conversion from template.FuncMap
type PlushFunctions template.FuncMap

// PlushOption object
type PlushOption struct {
	Functions	PlushFunctions
}

// PlushRender object
type PlushRender struct {
	option 	*PlushOption
	name 	string
	context gin.H
}

// NewPlushRender return the instance
func NewPlushRender(option *PlushOption) *PlushRender {
	return &PlushRender{
		option: option,
	}
}

// Instance return a new PlushRender struct per request
func (p PlushRender) Instance(name string, data interface{}) render.Render {
	return PlushRender{
		option : p.option,
		name   : name,
		context: data.(gin.H),
	}
}

/*
 * Implement for the render.Render
 */

// Render the template to the response
func (p PlushRender) Render(w http.ResponseWriter) error {
	assetBytes, err := view.Asset(string(p.name))
	if err != nil {
		return err
	}

	result, err := plush.BuffaloRenderer(string(assetBytes), p.context, p.option.Functions)
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
