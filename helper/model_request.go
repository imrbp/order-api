package helper

import (
	"restapi/model/entity"
	"restapi/model/web"
)

func ToCreateItemsEntity(items []web.ItemCreateRequest) []entity.Item {
	var itemResponse []entity.Item
	for _, data := range items {
		itemResponse = append(itemResponse, entity.Item{
			Code:        data.ItemCode,
			Description: data.Description,
			Quantity:    data.Quantity,
		})
	}
	return itemResponse
}

func ToCreateOrderEntity(order web.OrderCreateRequest) entity.Order {
	return entity.Order{
		OrderAt:      order.OrderAt,
		CustomerName: order.CustomerName,
		Items:        ToCreateItemsEntity(order.Items),
	}
}

func ToItemsUpdateEntity(items []web.ItemUpdateRequest) []entity.Item {
	var itemResponse []entity.Item
	for _, data := range items {
		itemResponse = append(itemResponse, entity.Item{
			Code:        data.ItemCode,
			Description: data.Description,
			Quantity:    data.Quantity,
		})
	}
	return itemResponse
}

func ToOrderUpdateEntity(order web.OrderUpdateRequest) entity.Order {
	return entity.Order{
		Id:           order.OrderId,
		OrderAt:      order.OrderAt,
		CustomerName: order.CustomerName,
		Items:        ToItemsUpdateEntity(order.Items),
	}
}
