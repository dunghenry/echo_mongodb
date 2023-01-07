package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// c.Request().Header.Add("token", "Bearer token")
		data := c.Request().Header.Get("token")
		if data != "" {
			return next(c)
		}
		fmt.Println("Not passed")
		return nil
	}
}
