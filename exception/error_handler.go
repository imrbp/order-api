package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi/helper"
	"restapi/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := context.Errors.Last()
		if notFoundError(context, err) {
			return
		}

		if validateError(context, err) {
			return
		}

		internalServiceError(context, err)
	}
}

func notFoundError(context *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		context.Set("Content-Type", "application/json")
		context.Status(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}
		helper.WriteToResponseBody(context.Writer, webResponse)
		return true
	} else {
		return false
	}
}

func validateError(context *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		context.Set("Content-Type", "application/json")
		context.Status(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}
		helper.WriteToResponseBody(context.Writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServiceError(context *gin.Context, err error) {
	context.Set("Content-Type", "application/json")
	context.Status(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}
	helper.WriteToResponseBody(context.Writer, webResponse)
}
