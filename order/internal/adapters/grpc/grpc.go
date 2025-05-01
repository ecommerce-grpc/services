package grpc

import (
	"context"

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
func (a Adapter) Create(cxt context.Context,
	request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	// TODO: Update order.proto
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
