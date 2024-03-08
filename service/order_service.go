package service

import (
	"context"
	"restapi/model/web"
)

type OrderService interface {
	Create(ctx context.Context, order web.OrderCreateRequest) web.OrderResponse
	Update(ctx context.Context, order web.OrderUpdateRequest) web.OrderResponse
	Delete(ctx context.Context, orderId int) web.OrderResponse
	GetById(ctx context.Context, orderId int) web.OrderResponse
	GetAll(ctx context.Context) []web.OrderResponse
}
