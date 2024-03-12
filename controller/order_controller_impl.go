package controller

import (
	"github.com/gin-gonic/gin"
	"restapi/helper"
	"restapi/model/web"
	"restapi/service"
	"strconv"
)

type OrderControllerImpl struct {
	Service service.OrderService
}

func NewOrderController(serviceOrder service.OrderService) OrderController {
	return &OrderControllerImpl{
		Service: serviceOrder,
	}
}

func (orderController OrderControllerImpl) Create(context *gin.Context) {
	orderCreateRequest := web.OrderCreateRequest{}
	helper.ReadFromRequestBody(context.Request, &orderCreateRequest)
	orderResult := orderController.Service.Create(context, orderCreateRequest)

	helper.WriteToResponseBody(context.Writer, orderResult)
}

func (orderController OrderControllerImpl) Update(context *gin.Context) {

	orderUpdateRequest := web.OrderUpdateRequest{}

	orderId := context.Param("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)
	helper.ReadFromRequestBody(context.Request, &orderUpdateRequest)
	orderUpdateRequest.OrderId = id
	orderResult := orderController.Service.Update(context, orderUpdateRequest)
	if orderResult.OrderId != id {
		helper.WriteToResponseBody(context.Writer, "Order Id Not Found")
	} else {
		helper.WriteToResponseBody(context.Writer, orderResult)
	}
}

func (orderController OrderControllerImpl) Delete(context *gin.Context) {

	orderId := context.Param("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderDeleted := orderController.Service.Delete(context, id)
	if orderDeleted.OrderId != id {
		helper.PanicIfError(err)
	}
	helper.WriteToResponseBody(context.Writer, "Success delete")
}

func (orderController OrderControllerImpl) FindById(context *gin.Context) {
	orderId := context.Param("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderFind := orderController.Service.GetById(context, id)

	helper.WriteToResponseBody(context.Writer, orderFind)
}

func (orderController OrderControllerImpl) FindAll(context *gin.Context) {

	ordersFind := orderController.Service.GetAll(context)

	helper.WriteToResponseBody(context.Writer, ordersFind)
}
