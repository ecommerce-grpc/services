package api

import (
	"context"

	"github.com/marcpires/grpc/ecommerce/payment/internal/application/core/domain"
	"github.com/marcpires/grpc/ecommerce/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db db.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	err := a.db.Save(ctx, &payment)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}