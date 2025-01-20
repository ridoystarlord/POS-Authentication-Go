package routes

import (
	"authentication/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(route fiber.Router) {
	route.Post("/login", controllers.Login)
	route.Post("/register", controllers.Register)
}