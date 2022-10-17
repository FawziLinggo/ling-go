package app

import (
	"FawziLinggo/ling-go/2.restful-api-go/controller"
	"FawziLinggo/ling-go/2.restful-api-go/execption"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.PanicHandler = execption.ErrorHandler

	return router
}
