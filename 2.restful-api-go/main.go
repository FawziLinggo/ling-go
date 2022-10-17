package main

import (
	"FawziLinggo/ling-go/2.restful-api-go/app"
	"FawziLinggo/ling-go/2.restful-api-go/controller"
	"FawziLinggo/ling-go/2.restful-api-go/helper"
	"FawziLinggo/ling-go/2.restful-api-go/middleware"
	"FawziLinggo/ling-go/2.restful-api-go/repository"
	"FawziLinggo/ling-go/2.restful-api-go/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "fawzi.linggo.id:3000",

		// NO AUTH
		// Handler: router,

		// AUTH
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
