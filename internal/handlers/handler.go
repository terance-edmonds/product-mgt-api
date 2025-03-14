package handlers

import (
	"fmt"
	"time"

	"github.com/terance-edmonds/product-mgt-api/internal/models"
	"github.com/terance-edmonds/product-mgt-api/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	products := storage.GetAllProducts()
	return c.JSON(models.Response{
		Status: "success",
		Data:   products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	currentTime := time.Now().UTC()
	product.ID = fmt.Sprintf("%d", currentTime.Unix())
	product.CreatedAt = currentTime
	product.UpdatedAt = currentTime

	storage.AddProduct(product)

	return c.JSON(models.Response{
		Status:  "success",
		Message: "Product created successfully",
		Data:    product,
	})
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("productId")
	if product, exists := storage.GetProduct(id); exists {
		return c.JSON(models.Response{
			Status: "success",
			Data:   product,
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("productId")
	if product, exists := storage.GetProduct(id); exists {
		var updatedProduct models.Product
		if err := c.BodyParser(&updatedProduct); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		updatedProduct.ID = id
		updatedProduct.CreatedAt = product.CreatedAt
		updatedProduct.UpdatedAt = time.Now().UTC()

		storage.UpdateProduct(id, updatedProduct)

		return c.JSON(models.Response{
			Status:  "success",
			Message: "Product updated successfully",
			Data:    updatedProduct,
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("productId")
	if storage.DeleteProduct(id) {
		return c.JSON(models.Response{
			Status:  "success",
			Message: "Product deleted successfully",
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
}

func AdjustInventory(c *fiber.Ctx) error {
	var adjustment models.InventoryAdjustment
	if err := c.BodyParser(&adjustment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if product, exists := storage.GetProduct(adjustment.ProductID); exists {
		product.StockQuantity += adjustment.QuantityAdjusted
		product.UpdatedAt = time.Now().UTC()
		storage.UpdateProduct(adjustment.ProductID, product)

		return c.JSON(models.Response{
			Status:  "success",
			Message: "Inventory adjusted successfully",
			Data: models.InventoryAdjustmentResponse{
				ProductID:        adjustment.ProductID,
				QuantityAdjusted: adjustment.QuantityAdjusted,
				NewStockQuantity: product.StockQuantity,
				Reason:           adjustment.Reason,
				AdjustmentDate:   time.Now().UTC(),
			},
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
}
