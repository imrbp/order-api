package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi/controller"
	"restapi/exception"
)

func OrderHandler(controller controller.OrderController) http.Handler {
	e := gin.New()

	e.Use(gin.Recovery())
	e.POST("/order", controller.Create)
	e.PUT("/order/:orderId", controller.Update)
	e.DELETE("/order/:orderId", controller.Delete)
	e.GET("/order", controller.FindAll)
	e.GET("/order/:orderId", controller.FindById)
	e.Use(exception.ErrorHandler())
	return e
}
