package order

import (
	"context"
	"fmt"

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

func (u *OrderRepoImpl) GetOrderByUserId(ctx context.Context, id uint) (result []order.Order, err error) {
	db := u.pgCln.GetClient()
	resultDb := db.Model(&order.Order{}).Preload("Item").Find(&result, "user_id = ?", id)
	if resultDb.Error != nil {
		return nil, resultDb.Error
	}
	return result, err
}

func (u *OrderRepoImpl) UpdateOrderById(ctx context.Context, id uint, updatedItem *order.Order) (err error) {
	var newOrder order.Order

	db := u.pgCln.GetClient()
	deleteItem := db.Where("order_id = ?", id).Delete(&order.Item{})
	if deleteItem.Error != nil {
		return deleteItem.Error
	}

	fmt.Println(deleteItem.Error)
	findOrder := db.Model(&order.Order{}).First(&newOrder, "id = ?", id)
	if findOrder.Error != nil {
		return findOrder.Error
	}

	newOrder.Item = updatedItem.Item
	db.Save(&newOrder)
	if err = db.Error; err != nil {
		return err
	}
	db.Model(&order.Order{}).First(&updatedItem, "id = ?", id)
	if err = db.Error; err != nil {
		return err
	}
	return err
}

func (u *OrderRepoImpl) DeleteOrderById(ctx context.Context, id uint) (err error) {
	db := u.pgCln.GetClient()
	deleteItem := db.Where("order_id = ?", id).Delete(&order.Item{})
	if deleteItem.Error != nil {
		return deleteItem.Error
	}

	deleteOrder := db.Where("id = ?", id).Delete(&order.Order{})
	if deleteOrder.Error != nil {
		return deleteOrder.Error
	}
	return err
}
