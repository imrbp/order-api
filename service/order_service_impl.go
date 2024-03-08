package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"restapi/exception"
	"restapi/helper"
	"restapi/model/web"
	"restapi/repository"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, db *gorm.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (service OrderServiceImpl) Create(ctx context.Context, order web.OrderCreateRequest) web.OrderResponse {
	errVal := service.Validate.Struct(order)
	helper.PanicIfError(errVal)

	orderResult, errSQL := service.OrderRepository.Create(ctx, helper.ToCreateOrderEntity(order))
	helper.PanicIfError(errSQL)

	return helper.ToOrderResponse(orderResult)
}

func (service OrderServiceImpl) Update(ctx context.Context, order web.OrderUpdateRequest) web.OrderResponse {
	errVal := service.Validate.Struct(order)
	helper.PanicIfError(errVal)

	orderFind, errFind := service.OrderRepository.GetById(ctx, order.OrderId)
	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}

	orderFind = helper.ToOrderUpdateEntity(order)
	orderUpdated, err := service.OrderRepository.Update(ctx, orderFind)
	helper.PanicIfError(err)

	return helper.ToOrderResponse(orderUpdated)
}

func (service OrderServiceImpl) Delete(ctx context.Context, orderId int) web.OrderResponse {
	orderFind, errFind := service.OrderRepository.GetById(ctx, orderId)

	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}
	errDeleted := service.OrderRepository.Delete(ctx, orderId)
	helper.PanicIfError(errDeleted)

	return helper.ToOrderResponse(orderFind)
}

func (service OrderServiceImpl) GetById(ctx context.Context, orderId int) web.OrderResponse {
	result, err := service.OrderRepository.GetById(ctx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderResponse(result)
}

func (service OrderServiceImpl) GetAll(ctx context.Context) []web.OrderResponse {
	allOrder := service.OrderRepository.GetAll(ctx)
	return helper.ToOrdersResponse(allOrder)
}
