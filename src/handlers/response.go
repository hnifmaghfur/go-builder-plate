package handlers

import "github.com/labstack/echo/v4"

func Response(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}
