package web

import (
	"time"
)

type OrderUpdateRequest struct {
	OrderId      int                 `json:"order_id" validate:"required"`
	CustomerName string              `json:"customer_name" validate:"required"`
	OrderAt      time.Time           `json:"order_at" validate:"required"`
	Items        []ItemUpdateRequest `json:"items" validate:"required"`
}

type ItemUpdateRequest struct {
	ItemId      int    `json:"item_id" validate:"required"`
	ItemCode    int    `json:"item_code" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	OrderId     int    `json:"order_id" validate:"required"`
}
