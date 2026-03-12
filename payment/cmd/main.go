package main

import (
	"github.com/marcpires/grpc/ecommerce/payment/config"
	"github.com/marcpires/grpc/ecommerce/payment/internal/adapters/db"
	"github.com/marcpires/grpc/ecommerce/payment/internal/adapters/grpc"
	"github.com/marcpires/grpc/ecommerce/payment/internal/application/core/api"

	log "github.com/sirupsen/logrus"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("error to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
