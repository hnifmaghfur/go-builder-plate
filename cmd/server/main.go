package main

import (
	"go-builder-plate/internal/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// new route instance
	route := echo.New()

	// cors middleware
	route.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// route handler
	route.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello World",
		})
	})

	// login handler
	route.POST("/login", handlers.Login)

	// start server
	route.Start(":3000")
}
