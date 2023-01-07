package routes

import (
	"trandung/api/controllers"
	"trandung/api/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {

	e.GET("/users", controllers.GetUsers, middleware.ServerHeader)
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	e.PUT("/users/:id", controllers.UpdateUser)
}
