package template

import (
	"errors"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"io"
)

type Template struct{}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	component, valid := data.(templ.Component)
	if !valid {
		return errors.New("data value is not valid templ component")
	}

	return component.Render(c.Request().Context(), w)
}

func NewTemplateRenderer(e *echo.Echo) {
	t := newTemplate()
	e.Renderer = t
}

func newTemplate() echo.Renderer {
	return &Template{}
}

func AssertRender(c echo.Context, statusCode int, component templ.Component) error {
	return c.Render(statusCode, "", component)
}

//NOTE: template object for text/html renderer
// type Template struct {
// 	Templates *template.Template
// }
//
// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.Templates.ExecuteTemplate(w, name, data)
// }
//
// func NewTemplateRenderer(e *echo.Echo, paths ...string) {
// 	tmpl := &template.Template{}
// 	for i := range paths {
// 		template.Must(tmpl.ParseGlob(paths[i]))
// 	}
// 	t := newTemplate(tmpl)
// 	e.Renderer = t
// }
//
// func newTemplate(templates *template.Template) echo.Renderer {
// 	return &Template{
// 		Templates: templates,
// 	}
// }
