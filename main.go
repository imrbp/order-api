package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"restapi/app"
	"restapi/controller"
	handler2 "restapi/handler"
	"restapi/helper"
	"restapi/repository"
	"restapi/service"
)

func main() {
	db := app.DBInit()
	//err := app.Seed()
	//if err != nil {
	//	panic(err)
	//}
	validate := validator.New()
	repositoryOrder := repository.NewOrderRepository(db)
	serviceOrder := service.NewOrderService(repositoryOrder, db, validate)
	controllerOrder := controller.NewOrderController(serviceOrder)
	handler := handler2.OrderHandler(controllerOrder)
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
