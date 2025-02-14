package services

import (
	"go-builder-plate/src/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var loginData models.LoginRequest
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request",
		})
	}

	if loginData.Username == "hanif" || loginData.Password == "123123" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request",
		})
	}
}
