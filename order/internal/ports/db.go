package ports

import (
	"context"

	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
)

// DBPort defines the funcionalities a type must implement to fulfill
type DBPort interface {
	Get(ctx context.Context, id int64) (domain.Order, error)
	Save(*domain.Order) error
}
