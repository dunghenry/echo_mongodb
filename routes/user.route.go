package routes

import (
	"trandung/api/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.GET("/", controllers.GetUsers)
}
