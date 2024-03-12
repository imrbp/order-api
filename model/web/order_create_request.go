package web

import (
	"time"
)

type OrderCreateRequest struct {
	CustomerName string              `json:"customerName" validate:"required,max=200,min=1"`
	OrderAt      time.Time           `json:"orderedAt" validate:"required"`
	Items        []ItemCreateRequest `json:"items"`
}

type ItemCreateRequest struct {
	ItemCode    string `json:"itemCode" validate:"required"`
	Description string `json:"description" validate:"required,max=200,min=1"`
	Quantity    int    `json:"quantity" validate:"required"`
}
