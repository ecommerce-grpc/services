package grpc

import (
	"context"
	"fmt"

	"github.com/ecommerce-grpc/payment"
	"github.com/marcpires/grpc/ecommerce/payment/internal/application/core/domain"
	"github.com/marcpires/grpc/ecommerce/payment/internal/ports"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	api ports.APIPort
}

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	log.WithContext(ctx).Info("Creating payment...")
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge: %v", err)).Err()
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}