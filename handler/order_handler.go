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
	e.POST("/orders", controller.Create)
	e.PUT("/orders/:orderId", controller.Update)
	e.DELETE("/orders/:orderId", controller.Delete)
	e.GET("/orders", controller.FindAll)
	e.GET("/orders/:orderId", controller.FindById)
	e.Use(exception.ErrorHandler())
	return e
}
