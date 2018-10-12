package render

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

    "github.com/zeuxisoo/go-tormag/pkg/view"
)

// BaseFunctions conversion from template.FuncMap
type BaseFunctions template.FuncMap

// BaseOption object
type BaseOption struct {
    Functions	BaseFunctions
}

// BaseRender object
type BaseRender struct {
    option 	*BaseOption
    name 	string
    data 	map[string]interface{}
}

// NewBaseRender return the instance
func NewBaseRender(option *BaseOption) *BaseRender {
    return &BaseRender{
        option: option,
    }
}

// Instance return a new BaseRender struct per request
func (b BaseRender) Instance(name string, data interface{}) render.Render {
    return BaseRender{
        option: b.option,
        name  : name,
        data  : data.(gin.H),
    }
}

/*
 * Implement for the render.Render
 */

// Render the template to the response
func (b BaseRender) Render(w http.ResponseWriter) error {
    assetBytes, err := view.Asset(string(b.name))
    if err != nil {
        return err
    }

    tpl, err := template.New("").Funcs((template.FuncMap)(b.option.Functions)).Parse(string(assetBytes))
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
        header["Content-Type"] = []string{ "text/html; charset=utf-8" }
    }
}
