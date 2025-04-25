package ports

import (
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
)

// PaymentPort defines the payment port behavior.
type PaymentPort interface {
	Charge(*domain.Order) error
}
