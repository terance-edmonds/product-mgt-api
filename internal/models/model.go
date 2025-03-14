package models

import "time"

type Product struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	Category      string    `json:"category"`
	StockQuantity int       `json:"stock_quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type InventoryAdjustment struct {
	ProductID        string `json:"product_id"`
	QuantityAdjusted int    `json:"quantity_adjusted"`
	Reason           string `json:"reason"`
}

type InventoryAdjustmentResponse struct {
	ProductID        string    `json:"product_id"`
	QuantityAdjusted int       `json:"quantity_adjusted"`
	NewStockQuantity int       `json:"new_stock_quantity"`
	Reason           string    `json:"reason"`
	AdjustmentDate   time.Time `json:"adjustment_date"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
