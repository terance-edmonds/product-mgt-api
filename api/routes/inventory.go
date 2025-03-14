package routes

import (
	"github.com/terance-edmonds/product-mgt-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupInventoryRoutes(app fiber.Router) {
	app.Post("/inventory/adjustment", handlers.AdjustInventory)
}
