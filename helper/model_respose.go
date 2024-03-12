package helper

import (
	"restapi/model/entity"
	"restapi/model/web"
)

func ToItemsResponse(items []entity.Item) []web.ItemResponse {
	var itemResponse []web.ItemResponse
	for _, data := range items {
		itemResponse = append(itemResponse, web.ItemResponse{
			ItemCode:    data.Code,
			Description: data.Description,
			Quantity:    data.Quantity,
		})
	}
	return itemResponse
}

func ToOrderResponse(order entity.Order) web.OrderResponse {
	return web.OrderResponse{
		OrderId:      order.Id,
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
