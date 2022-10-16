package order

import (
	"context"
)

type OrderUsecase interface {
	CreateOrderSvc(ctx context.Context, input Order) (result Order, err error)
}
