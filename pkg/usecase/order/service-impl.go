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

func (u *OrderUsecaseImpl) GetOrderByUserIdSvc(ctx context.Context, id uint) (result []order.Order, err error) {

	result, err = u.orderRepo.GetOrderByUserId(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (u *OrderUsecaseImpl) UpdateOrderByIdSvc(ctx context.Context, id uint, updatedOrder order.Order) (result order.Order, err error) {

	err = u.orderRepo.UpdateOrderById(ctx, id, &updatedOrder)

	if err != nil {
		return result, err
	}

	return updatedOrder, err
}

func (u *OrderUsecaseImpl) DeleteOrderByIdSvc(ctx context.Context, id uint) (err error) {

	err = u.orderRepo.DeleteOrderById(ctx, id)

	if err != nil {
		return err
	}

	return err
}
