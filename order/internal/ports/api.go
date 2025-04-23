// Package ports defines the APIPort
package ports

import (
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
)

// APIPorter defined the port interface a type must satisfy to implement it.
type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}