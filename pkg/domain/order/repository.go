package order

import "context"

type OrderRepo interface {
	CreateOrder(ctx context.Context, input *Order) (err error)
}
