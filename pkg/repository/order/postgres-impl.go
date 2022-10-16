package order

import (
	"context"

	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/config/postgres"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/order"
)

type OrderRepoImpl struct {
	pgCln postgres.PostgreClient
}

func NewOrderRepo(pgcln postgres.PostgreClient) order.OrderRepo {
	return &OrderRepoImpl{pgCln: pgcln}
}

func (u *OrderRepoImpl) CreateOrder(ctx context.Context, input *order.Order) (err error) {
	db := u.pgCln.GetClient()
	result := db.Model(&order.Order{}).Create(&input)
	if err := result.Error; err != nil {
		return err
	}
	return err
}
