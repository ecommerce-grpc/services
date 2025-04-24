package grpc

import (
	"context"

	"github.com/marcpires/grpc/ecommerce/services/order" // service stub
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/domain"
)

func (a Adapter) Create(cxt context.Context, 
	request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
		var orderItems []domain.OrderItem
		for _, orderItem := range request.OrderItems {
			orderItem = append(orderItems, domain.OrderItem{
				ProductCode: orderItem.ProductCode,
				UnitPrice: orderItem.UnitPrice,
				Quantity: orderItem.Quantity,
			})
		}

		newOrder := domain.NewOrder(request.UserId, orderItems)
		result, err := a.api.PlaceOrder(newOrder)
		if err != nil {
			return nil, err
		}

		return &order.CreateOrderResponse{OrderId: result.ID}, nil
	}
