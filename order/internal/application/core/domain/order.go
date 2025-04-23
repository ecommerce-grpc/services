package domain

import (
	"time"
)

// OrderItem is a domain object representing an Order Item
type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

// Order is a domain object representing a customer order
type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreateAt   int64       `json:"created_at"`
}

// NewOrder returns the items of customer order
func NewOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		CreateAt:   time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}
