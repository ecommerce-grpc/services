package ports

import (
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
)

// DBPort defines the funcionalities a type must implement to fulfill
type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}