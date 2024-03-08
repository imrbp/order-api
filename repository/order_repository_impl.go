package repository

import (
	"context"
	"restapi/model/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{DB: db}
}

func (repository *OrderRepositoryImpl) Create(ctx context.Context, order entity.Order) (entity.Order, error) {
	data := entity.Order{
		CustomerName: order.CustomerName,
		OrderAt:      order.OrderAt,
		Items:        order.Items,
	}
	result := repository.DB.WithContext(ctx).Create(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}
func (repository *OrderRepositoryImpl) Update(ctx context.Context, order entity.Order) (entity.Order, error) {
	result := repository.DB.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)
	if result.Error != nil {
		return order, result.Error
	}
	return order, nil
}
func (repository *OrderRepositoryImpl) Delete(ctx context.Context, orderId int) error {
	orderDeleting, err := repository.GetById(ctx, orderId)
	if err != nil {
		return err
	}
	result := repository.DB.WithContext(ctx).Select(clause.Associations).Delete(orderDeleting)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (repository *OrderRepositoryImpl) GetById(ctx context.Context, orderId int) (entity.Order, error) {
	order := entity.Order{}
	findOrder := repository.DB.WithContext(ctx).Preload(clause.Associations).Model(&entity.Order{}).First(&order, orderId).Scan(&order)
	if findOrder.Error != nil {
		return order, findOrder.Error
	}
	return order, nil

}
func (repository *OrderRepositoryImpl) GetAll(ctx context.Context) []entity.Order {
	orders := &[]entity.Order{}
	repository.DB.WithContext(ctx).Preload(clause.Associations).Model(&entity.Order{}).Find(&orders)

	return *orders
}
