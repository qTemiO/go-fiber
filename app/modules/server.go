package modules

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/template/django/v3"

	swagger "github.com/gofiber/contrib/swagger"

	"app/config"
)

// NewServer creates a new Fiber app and sets up the routes
func NewServer(context context.Context, config *config.Config) *fiber.App {
	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "Gorag",
		Views:        engine,
	})
	app.Use(swagger.New(swagger.Config{
		FilePath: "./docs/swagger.json",
	}))
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        5 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
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

	api := app.Group("/api")
	v1 := api.Group("/v1")

	fmt.Printf("Enter other cool router %s", v1)

	return app
}
