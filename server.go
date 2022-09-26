package main

import (
	"trandung/api/configs"
	"trandung/api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	configs.ConnectDB()
	routes.UserRoute(e)
	e.Logger.Fatal(e.Start("localhost:8080"))
}
