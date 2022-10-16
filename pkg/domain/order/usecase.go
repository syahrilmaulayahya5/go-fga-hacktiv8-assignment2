package order

import (
	"context"
)

type OrderUsecase interface {
	CreateOrderSvc(ctx context.Context, input Order) (result Order, err error)
	GetOrderByUserIdSvc(ctx context.Context, id uint) (result []Order, err error)
	UpdateOrderByIdSvc(ctx context.Context, id uint, updatedOrder Order) (result Order, err error)
	DeleteOrderByIdSvc(ctx context.Context, id uint) (err error)
}
