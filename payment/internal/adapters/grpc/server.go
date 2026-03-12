package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/ecommerce-grpc/payment"
	"github.com/marcpires/grpc/ecommerce/payment/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run executes the gRPC server. Can be moved to a ecommerce std library.
func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	payment.RegisterPaymentServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if sListenErr := grpcServer.Serve(listen); sListenErr != nil {
		log.Fatalf("Failed to serve grpc on port ")
	}
}
