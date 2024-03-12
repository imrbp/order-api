package web

import (
	"time"
)

type OrderUpdateRequest struct {
	OrderId      int                 `json:"id" validate:"required"`
	CustomerName string              `json:"customerName" validate:"required"`
	OrderAt      time.Time           `json:"orderedAt" validate:"required"`
	Items        []ItemUpdateRequest `json:"items" validate:"required"`
}

type ItemUpdateRequest struct {
	ItemCode    string `json:"itemCode" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}
