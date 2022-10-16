package user

import (
	"context"

	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/order"
)

type OrderUsecaseImpl struct {
	orderRepo order.OrderRepo
}

func NewOrderUsecase(orderRepo order.OrderRepo) order.OrderUsecase {
	return &OrderUsecaseImpl{orderRepo: orderRepo}

}

func (u *OrderUsecaseImpl) CreateOrderSvc(ctx context.Context, input order.Order) (result order.Order, err error) {

	err = u.orderRepo.CreateOrder(ctx, &input)

	if err != nil {
		return result, err
	}
	return input, err
}
