// main is the program package to run the Order service
// core application needs a DB adapters and gRPC server needs the core application as dependencies.
package main

import (
	"log"

	"github.com/marcpires/grpc/ecommerce/order/config"
	"github.com/marcpires/grpc/ecommerce/order/internal/adapters/db"
	"github.com/marcpires/grpc/ecommerce/order/internal/adapters/grpc"
	"github.com/marcpires/grpc/ecommerce/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	app := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(app, config.GetApplicationPort())
	grpcAdapter.Run()
}
