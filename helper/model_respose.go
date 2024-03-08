package helper

import (
	"restapi/model/entity"
	"restapi/model/web"
)

func ToItemsResponse(items []entity.Item) []web.ItemResponse {
	var itemResponse []web.ItemResponse
	for _, data := range items {
		itemResponse = append(itemResponse, web.ItemResponse{
			ItemId:      data.ItemId,
			ItemCode:    data.ItemCode,
			Description: data.Description,
			Quantity:    data.Quantity,
			OrderId:     data.OrderId,
		})
	}
	return itemResponse
}

func ToOrderResponse(order entity.Order) web.OrderResponse {
	return web.OrderResponse{
		OrderId:      order.OrderId,
		OrderAt:      order.OrderAt,
		CustomerName: order.CustomerName,
		Items:        ToItemsResponse(order.Items),
	}
}

func ToOrdersResponse(orders []entity.Order) []web.OrderResponse {
	var responseOrder []web.OrderResponse
	for _, item := range orders {
		responseOrder = append(responseOrder, ToOrderResponse(item))
	}
	return responseOrder
}
