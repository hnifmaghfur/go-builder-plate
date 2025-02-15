package handlers

import (
	"fmt"
	"go-builder-plate/src/models"
	"go-builder-plate/src/services"
	"go-builder-plate/src/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

// this handler is used for validate reqeust data

func Login(c echo.Context) error {
	var loginData models.LoginRequest

	// bind request body
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request bind",
		})
	}

	// validate request body
	if err := validation.ValidateLogin(loginData); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request validate",
		})
	}

	status, data := services.Login(loginData)
	return c.JSON(status, data)
}
