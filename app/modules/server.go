package modules

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/template/django/v3"

	swagger "github.com/gofiber/contrib/swagger"

	"app/config"
)

// NewServer creates a new Fiber app and sets up the routes
func NewServer(context context.Context, config *config.Config) *fiber.App {
	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(swagger.New(swagger.Config{
		FilePath: "./docs/swagger.json",
	}))

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))
	app.Static("/static", "./static")
	app.Get("/", func(context *fiber.Ctx) error {
		// Render with and extends
		return context.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	// api := app.Group("/api")
	// v1 := api.Group("/v1")

	// creatorRouter := v1.Group("/creator")
	// creator.SetupRoutes(creatorRouter, config)

	// validatorRouter := v1.Group("/validator")
	// validator.SetupRoutes(validatorRouter, config)
	return app
}
