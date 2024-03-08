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

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResult,
	}
	helper.WriteToResponseBody(context.Writer, webResponse)
}

func (orderController OrderControllerImpl) Update(context *gin.Context) {

	orderUpdateRequest := web.OrderUpdateRequest{}

	orderId := context.Param("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	helper.ReadFromRequestBody(context.Request, &orderUpdateRequest)
	orderResult := orderController.Service.Update(context, orderUpdateRequest)
	webResponse := web.WebResponse{}
	if orderResult.OrderId != id {
		webResponse = web.WebResponse{
			Code:   200,
			Status: "Order Id Not Found",
			Data:   nil,
		}
	} else {
		webResponse = web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   orderResult,
		}
	}
	helper.WriteToResponseBody(context.Writer, webResponse)
}

func (orderController OrderControllerImpl) Delete(context *gin.Context) {

	orderId := context.Param("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderDeleted := orderController.Service.Delete(context, id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderDeleted,
	}

	helper.WriteToResponseBody(context.Writer, webResponse)
}

func (orderController OrderControllerImpl) FindById(context *gin.Context) {
	orderId := context.Param("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderFind := orderController.Service.GetById(context, id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderFind,
	}

	helper.WriteToResponseBody(context.Writer, webResponse)
}

func (orderController OrderControllerImpl) FindAll(context *gin.Context) {

	ordersFind := orderController.Service.GetAll(context)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ordersFind,
	}

	helper.WriteToResponseBody(context.Writer, webResponse)
}
