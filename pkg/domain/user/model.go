package user

import "github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/order"

type User struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	FirstName string        `json:"first_name" gorm:"column:first_name"`
	LastName  string        `json:"last_name" gorm:"column:last_name"`
	Email     string        `json:"email" gorm:"column:email;unique"`
	Orders    []order.Order `json:"-"`
}
