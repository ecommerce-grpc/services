// Package ports defines the APIPort
package ports

import (
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
)

// APIPorter defines the API port behavior.
type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}