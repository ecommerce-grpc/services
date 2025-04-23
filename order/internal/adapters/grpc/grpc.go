package grpc

import "github.com/marcpires/grpc/ecommerce/order/internal/ports"

type Adapter struct {
	api ports.APIPort //Core application dependency
	port int
	order.UnimplementedOrderServer //forward compatibility support
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}