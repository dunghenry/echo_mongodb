package routes

import (
	"trandung/api/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.GET("/", controllers.GetUsers)
	e.POST("/", controllers.CreateUser)
	e.GET("/:id", controllers.GetUser)
	e.DELETE("/:id", controllers.DeleteUser)
	e.PUT("/:id", controllers.UpdateUser)
}
