package routes

import (
	"authentication/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api:=app.Group("/api/v1")
	SetupAuthRoutes(api.Group("/auth"))
	SetupBookRoutes(api.Group("/book",middleware.IsAuthenticated))
}