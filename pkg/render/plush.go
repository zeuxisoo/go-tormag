package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/gobuffalo/plush"

	"github.com/zeuxisoo/go-tormag/pkg/view"
)

// PlushRender object
type PlushRender struct {
	engine  *gin.Engine
	name 	string
	context *plush.Context
}

// NewPlushRender return the instance
func NewPlushRender(engine *gin.Engine) *PlushRender {
	return &PlushRender{
		engine: engine,
	}
}

// Instance return a new PlushRender struct per request
func (p PlushRender) Instance(name string, data interface{}) render.Render {
	return PlushRender{
		engine : p.engine,
		name   : name,
		context: plush.NewContextWith(data.(gin.H)),
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

	for name, callback := range p.engine.FuncMap {
		p.context.Set(name, callback)
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
