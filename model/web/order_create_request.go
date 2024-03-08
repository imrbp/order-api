package web

import (
	"time"
)

type OrderCreateRequest struct {
	CustomerName string              `json:"customer_name" validate:"required,max=200,min=1"`
	OrderAt      time.Time           `json:"order_at" validate:"required"`
	Items        []ItemCreateRequest `json:"items"`
}

type ItemCreateRequest struct {
	ItemCode    int    `json:"item_code" validate:"required"`
	Description string `json:"description" validate:"required,max=200,min=1"`
	Quantity    int    `json:"quantity" validate:"required"`
}
