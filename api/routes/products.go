package routes

import (
	"github.com/terance-edmonds/product-mgt-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App) {
	app.Get("/products", handlers.GetProducts)
	app.Post("/products", handlers.CreateProduct)
	app.Get("/products/:productId", handlers.GetProduct)
	app.Put("/products/:productId", handlers.UpdateProduct)
	app.Delete("/products/:productId", handlers.DeleteProduct)
}
