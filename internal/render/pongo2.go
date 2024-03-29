package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"

	"github.com/zeuxisoo/go-tormag/views"

	"gopkg.in/flosch/pongo2.v3"
)

// Pongo2FilterFunctions conversion from pongo2.FilterFunction
type Pongo2FilterFunctions map[string]pongo2.FilterFunction

// Pongo2Option object
type Pongo2Option struct {
    Filters	Pongo2FilterFunctions
}

// Pongo2Render object
type Pongo2Render struct {
    option	*Pongo2Option
    name	string
    data	gin.H
}

// NewPongo2Render return the instance
func NewPongo2Render(option *Pongo2Option) *Pongo2Render {
    for name, filter := range option.Filters {
        pongo2.RegisterFilter(name, filter)
    }

    return &Pongo2Render{
        option: option,
    }
}

// Instance return a new Pongo2Render struct per request
func (p Pongo2Render) Instance(name string, data interface{}) render.Render {
    return Pongo2Render{
        option: p.option,
        name  : name,
        data  : data.(gin.H),
    }
}

/*
 * Implement for the render.Render
 */

// Render the template to the response
func (p Pongo2Render) Render(w http.ResponseWriter) error {
    assetBytes, err := views.Files.ReadFile(p.name)
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
