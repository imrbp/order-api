package web

import (
	"time"
)

type OrderResponse struct {
	OrderId      int            `json:"order_id"`
	CustomerName string         `json:"customer_name"`
	OrderAt      time.Time      `json:"order_at"`
	Items        []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ItemId      int    `json:"item_id"`
	ItemCode    int    `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}
