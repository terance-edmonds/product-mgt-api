package storage

import (
	"github.com/terance-edmonds/product-mgt-api/internal/config"
	"github.com/terance-edmonds/product-mgt-api/internal/models"

	"github.com/google/uuid"
)

var products = make(map[string]models.Product)

func LoadInitialData() {
	// Read initial data from JSON file
	initialData := config.LoadInitialData()
	initialProducts := initialData.Products

	// Populate the products map
	for _, product := range initialProducts {
		products[product.ID] = product
	}
}

func GetAllProducts() []models.Product {
	var productList []models.Product
	for _, product := range products {
		productList = append(productList, product)
	}
	return productList
}

func GetProduct(id string) (models.Product, bool) {
	product, exists := products[id]
	return product, exists
}

func AddProduct(product models.Product) string {
	if product.ID == "" {
		// Generate a new UUID if ID is not provided
		product.ID = uuid.New().String()
	}
	products[product.ID] = product
	return product.ID
}

func UpdateProduct(id string, product models.Product) {
	products[id] = product
}

func DeleteProduct(id string) bool {
	if _, exists := products[id]; exists {
		delete(products, id)
		return true
	}
	return false
}
