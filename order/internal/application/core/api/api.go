package api

import (
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
	"github.com/marcpires/grpc/ecommerce/order/internal/ports"
)

// Application defines the application.
type Application struct {
	db ports.DBPort
}

// NewApplication return an Application pointer with a database dependency.
func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

// PlaceOrder creates a new Order. Implements the APIPorter interface
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

