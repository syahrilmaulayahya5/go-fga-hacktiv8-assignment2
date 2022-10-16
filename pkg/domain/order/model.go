package order

import (
	"time"
)

type Order struct {
	CreatedAt time.Time `json:"orderedAt"`
	ID        uint      `json:"-" gorm:"primaryKey"`
	UserID    uint      `json:"userId"`
	Item      []Item    `json:"items"`
}

type Item struct {
	ID          uint   `json:"-"`
	Code        string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}
