package execption

import (
	"FawziLinggo/ling-go/2.restful-api-go/helper"
	"FawziLinggo/ling-go/2.restful-api-go/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}
	if validationErros(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)

}

func validationErros(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	execption, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD Requeest",
			Data:   execption.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	execption, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   execption.Error,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
