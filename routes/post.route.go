package routes

import (
	"trandung/api/controllers"

	"github.com/labstack/echo/v4"
)

func PostRoute(e *echo.Echo) {
	e.GET("/posts", controllers.GetPosts)
	e.POST("/posts", controllers.CreatePost)
	e.GET("/posts/:id", controllers.GetPost)
	e.PUT("/posts/:id", controllers.UpdatePost)
	e.DELETE("/posts/:id", controllers.DeletePost)
}
