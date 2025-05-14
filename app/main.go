package main

import (
	"app/config"
	"app/modules"
	"context"
	"fmt"
	"log"
)

// @title License Service API
// @version 1.0
// @description API for creating licenses

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	config := config.LoadConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := modules.NewServer(ctx, config)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", *config.Host, *config.Port)))
}
