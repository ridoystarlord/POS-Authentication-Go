package routes

import (
	"authentication/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(route fiber.Router) {
	route.Post("/new", controllers.CreateBook)
	route.Delete("/:id", controllers.DeleteBook)
	route.Get("/",  controllers.GetBook)
	route.Get("/:id", controllers.GetBookById)
}