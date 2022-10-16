package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"trandung/api/configs"
	"trandung/api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	configs.ConnectDB()
	e.Static("/public", "public")
	routes.UserRoute(e)
	routes.PostRoute(e)
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	// views
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"title":    "Home Page!",
			"fullName": "TranVanDung",
		})
	})

	//Upload file

	e.POST("/upload", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		file, err := c.FormFile("file")

		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		data := time.Now().Unix()
		filename := strconv.Itoa(int(data)) + "." + strings.Split(file.Filename, ".")[1]
		dst, err := os.Create("./public/images/" + filename)
		if err != nil {
			return err
		}

		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return err
		}
		return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
	})

	e.Renderer = renderer
	e.Logger.Fatal(e.Start(":8080"))
}
