package handlers

import (
	"go-builder-plate/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// this handler is used for validate reqeust data

func Login(c echo.Context) error {
	result := services.Login(c)
	return c.JSON(http.StatusOK, result)
}
