package grpc

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/ecommerce-grpc/order" // service stub
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
	"github.com/marcpires/grpc/ecommerce/order/internal/ports"
)

type Adapter struct {
	api                            ports.APIPort //Core application dependency
	port                           int
	order.UnimplementedOrderServer //forward compatibility support
}

// NewAdapter returns a gRPC adapter
func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

// Create handles the Order request endpoint
func (a Adapter) Create(ctx context.Context,
	request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	log.WithContext(ctx).Info("Making an order")

	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
