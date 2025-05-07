package payment

import (
	"context"
	"time"

	"github.com/ecommerce-grpc/payment"
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

// Adapter defines the payment adapter that handles the Payment operations
type Adapter struct {
	payment payment.PaymentClient
}

// unaryInterceptorRetry returns a retry pattern using Linar Backoff
func unaryInterceptorRetry() grpc.DialOption {
	return grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
	))
}

// NewAdapter returns a new payment Adapter using disabled TLS credentials or an error if applicable.
func NewAdapter(paymentServiceURL string) (*Adapter, error) {
	var opts []grpc.DialOption

	// add interceptors for unary connections
	// opts = append(opts, unaryInterceptorRetry())

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(paymentServiceURL, opts...)
	if err != nil {
		return nil, err
	}
	// This caused an error as the connection is get closed before the service is able to servie an Order.
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
