package views

//                                      +------------------------------------------------+
//                                      |                                                |
//                                      |  Instance +----> BaseRender +----> Render      |
//                                      |                                      ^         |
// Gin.HTMLRender +--> BaseRender +--+  |                    (new)             |         |
//                                      |                                      +         |
//                    (*NewRender)      |                               WriteContentType |
//                                      |                                                |
//                                      +------------------------------------------------+

import (
	"bytes"
	"net/http"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

// BaseRender object
type BaseRender struct {
	name string
	data map[string]interface{}
}

// NewRender return the instance
func NewRender() *BaseRender {
	return &BaseRender{}
}

// Instance return a new BaseRender struct per request
func (b BaseRender) Instance(name string, data interface{}) render.Render {
	return BaseRender{
		name: name,
		data: data.(gin.H),
	}
}

/*
 * Implement for the render.Render
 */

// Render the template to the response
func (b BaseRender) Render(w http.ResponseWriter) error {
	assetBytes, err := Asset(string(b.name))
	if err != nil {
		return err
	}

	tpl, err := template.New(b.name).Parse(string(assetBytes))
	if err != nil {
		return err
	}

	var result bytes.Buffer
	if err = tpl.Execute(&result, b.data); err != nil {
		return err
	}

	b.WriteContentType(w)

	w.Write([]byte(result.String()))

	return nil
}

// WriteContentType add the content type header to the response
func (b BaseRender) WriteContentType(w http.ResponseWriter) {
	header := w.Header()

	if contentType := header["Content-Type"]; len(contentType) == 0 {
		header["Content-Type"] = []string{
			"text/html",
			"charset=utf-8",
		}
	}
}
