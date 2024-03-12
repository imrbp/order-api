package web

import (
	"time"
)

type OrderResponse struct {
	OrderId      int            `json:"id"`
	CustomerName string         `json:"customerName"`
	OrderAt      time.Time      `json:"orderAt"`
	Items        []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
