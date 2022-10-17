package controller

import (
	"net/http"
	"strconv"

	"FawziLinggo/ling-go/2.restful-api-go/helper"
	"FawziLinggo/ling-go/2.restful-api-go/model/web"
	"FawziLinggo/ling-go/2.restful-api-go/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)

}
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoriesResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
