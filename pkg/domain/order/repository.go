package order

import "context"

type OrderRepo interface {
	GetOrderByUserId(ctx context.Context, id uint) (result []Order, err error)
	CreateOrder(ctx context.Context, input *Order) (err error)
	UpdateOrderById(ctx context.Context, id uint, updatedOrder *Order) (err error)
	DeleteOrderById(ctx context.Context, id uint) (err error)
}
