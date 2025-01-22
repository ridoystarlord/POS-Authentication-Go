package main

import (
	_ "authentication/docs"
	"authentication/routes"
	"authentication/storage"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	"github.com/ilyakaznacheev/cleanenv"
)

// @title						Book API
// @version					1.0
// @description				This is a sample swagger for simple book api
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @host						localhost:8000
// @BasePath					/api/v1
func main() {
	var config storage.Config

	// Load configuration
	if err := cleanenv.ReadConfig(".env", &config); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// Connect to the database
	storage.ConnectDatabase(&config)
	storage.MigrateDB(storage.DB)
	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(compress.New())
	app.Use(healthcheck.New())
	// app.Use(limiter.New())
	// app.Use(logger.New())
	app.Get("/metrics", monitor.New())
	app.Static("/", "./public")
	app.Get("/api-docs/*", swagger.HandlerDefault)

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	routes.SetupRoutes(app)
	fmt.Println("Server running on port 8000")
	app.Listen(":8000")
}
