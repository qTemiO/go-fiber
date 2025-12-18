package main

import (
	"app/config"
	"app/modules"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/ansrivas/fiberprometheus/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/lmittmann/tint"
)

func Setup() (*fiber.App, *config.Config) {
	w := os.Stderr
	// Create a new logger
	logger := slog.New(tint.NewHandler(w, nil))

	// Set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	slog.Info("Logger set", "logger", logger)

	config := config.LoadConfig("", "")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := modules.NewServer(ctx, config)

	// Prometheus - metrics
	prometheus := fiberprometheus.New("GORAG mertics-prometheus")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)
	// Fiber - monitor
	app.Get("/monitor", monitor.New(monitor.Config{Title: "GORAG metrics-monitor"}))

	return app, config
}

// @title Golang Fiber template Service
// @version 0.1.0
// @description API for education and myself
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	app, config := Setup()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", *config.Host, *config.Port)))
}
