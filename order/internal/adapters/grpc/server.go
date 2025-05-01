package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/ecommerce-grpc/order" // service stub
	"github.com/marcpires/grpc/ecommerce/order/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


// Run executes an gRPC server instance on a given port
// in development mode it runs with reflecion enabled.
func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if sListenErr := grpcServer.Serve(listen); sListenErr != nil {
		log.Fatalf("Failed to serve grpc on port ")
	}
}
