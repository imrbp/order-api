package repository

import (
	"context"
	"restapi/model/entity"
)

type OrderRepository interface {
	Create(ctx context.Context, order entity.Order) (entity.Order, error)
	Update(ctx context.Context, order entity.Order) (entity.Order, error)
	Delete(ctx context.Context, orderId int) error
	GetById(ctx context.Context, orderId int) (entity.Order, error)
	GetAll(ctx context.Context) []entity.Order
}
