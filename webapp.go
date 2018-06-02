package main

import (
	"github.com/labstack/echo"
	"AriaCTFer/handler"
	"html/template"
	"io"
)

/*テンプレートエンジン関連*/
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	e := echo.New()
	e.Static("/static", "assets")

	e.Renderer = &TemplateRenderer{
		template.Must(template.ParseGlob("template/*.html")),
	}
	e.GET("/", handler.IndexPage())
	e.GET("/register", handler.Register_GET_Page())
	e.POST("/register", handler.Register_POST_Page())

	e.Logger.Fatal(e.Start(":8080"))
}
