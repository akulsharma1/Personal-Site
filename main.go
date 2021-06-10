package main

import (
	"html/template"
	"io"
	"github.com/labstack/echo"
	"personalsite/handler"
)

type TemplateRegistry struct {
	templates *template.Template
}
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	e := echo.New()
	e.Renderer = &TemplateRegistry {
		templates: template.Must(template.ParseGlob("view/*.html")),
	}

	e.GET("/", handler.HomeHandler)
	e.Logger.Fatal(e.Start(":3000"))
}