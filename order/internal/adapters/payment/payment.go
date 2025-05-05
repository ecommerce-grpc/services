package payment

import (
	"context"

	"github.com/ecommerce-grpc/payment"
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Adapter defines the payment adapter that handles the Payment operations
type Adapter struct {
	payment payment.PaymentClient
}

// NewAdapter returns a new payment Adapter using disabled TLS credentials or an error if applicable.
func NewAdapter(paymentServiceURL string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(paymentServiceURL, opts...)
	if err != nil {
		return nil, err
	}
	// This causes an error as the connection is closed before the Order request is made
	// defer conn.Close()
	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}

// Charge handles the payment operation, implements the payment port using the payment stub
func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}
