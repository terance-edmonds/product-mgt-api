package main

import (
	"fmt"
	"time"

	"github.com/terance-edmonds/product-mgt-api/api/routes"
	"github.com/terance-edmonds/product-mgt-api/internal/config"
	"github.com/terance-edmonds/product-mgt-api/internal/storage"
	"github.com/terance-edmonds/product-mgt-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:               "choreo-product-mgt-api",
		ReadTimeout:           time.Second * 2,
		Prefork:               false,
		DisableKeepalive:      true,
		DisableStartupMessage: true,
		ErrorHandler:          utils.FiberErrorHandler,
	})

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("fiber error: %v", err)
	}

	// Load initial data
	storage.LoadInitialData()

	// Create a group for the base path
	api := app.Group(cfg.BasePath)

	// Setup routes
	routes.SetupProductRoutes(api)
	routes.SetupInventoryRoutes(api)

	// Print startup message before starting the server
	fmt.Printf("API Running on port: %v\n", cfg.Port)

	// Start the server
	if err := app.Listen(fmt.Sprintf("%v:%v", cfg.Hostname, cfg.Port)); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
