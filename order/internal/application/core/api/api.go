package api

import (
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
	"github.com/marcpires/grpc/ecommerce/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Application defines the application.
type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

// NewApplication return an Application pointer with a database dependency.
func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
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
	
	paymentErr := a.payment.Charge(&order)
	// Better error handling with more detailed errors
	if paymentErr != nil {
		st, _ := status.FromError(paymentErr)
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field: "payment",
			Description: st.Message(),
			Reason: "PAYMENT_INVALID_ARGUMENT",
		}
		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetails, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusWithDetails.Err()
	}
	return order, nil
}
