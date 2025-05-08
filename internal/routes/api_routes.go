package routes

import (
	userHandler "go-builder-plate/internal/handlers/http/v1/user"

	"github.com/labstack/echo/v4"
)

// SetupRoutes configures all the API routes
func SetupRoutes(e *echo.Echo) {

	// Define your API routes here

	// user handler
	api := e.Group("/api/v1")
	{
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", userHandler.Login)
		}
	}
}
