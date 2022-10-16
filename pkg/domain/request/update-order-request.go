package request

import "github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/order"

type UpdateOrderRequest struct {
	Item []order.Item `json:"items"`
}

func ToDomain(input UpdateOrderRequest) (result order.Order) {

	result.Item = input.Item
	return result
}
