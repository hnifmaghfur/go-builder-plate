package auth

import (
	"github.com/labstack/echo/v4"
)

func JWT(text string) string {
	e.Use(echo.JWT([]byte("secret")))
	return text
}
