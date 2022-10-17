package main

import (
	"FawziLinggo/ling-go/2.restful-api-go/app"
	"FawziLinggo/ling-go/2.restful-api-go/controller"
	"FawziLinggo/ling-go/2.restful-api-go/execption"
	"FawziLinggo/ling-go/2.restful-api-go/helper"
	"FawziLinggo/ling-go/2.restful-api-go/repository"
	"FawziLinggo/ling-go/2.restful-api-go/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = execption.ErrorHandler

	server := http.Server{
		Addr:    "fawzi.linggo.id:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
