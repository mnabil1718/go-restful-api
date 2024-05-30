package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/mnabil1718/go-restful-api/app"
	"github.com/mnabil1718/go-restful-api/controller"
	"github.com/mnabil1718/go-restful-api/exception"
	"github.com/mnabil1718/go-restful-api/helper"
	"github.com/mnabil1718/go-restful-api/repository"
	"github.com/mnabil1718/go-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	apiV1RootPath := "/api/v1"

	router.GET(apiV1RootPath+"/categories", categoryController.FindAll)
	router.POST(apiV1RootPath+"/categories", categoryController.Create)
	router.GET(apiV1RootPath+"/categories/:categoryId", categoryController.FindById)
	router.PUT(apiV1RootPath+"/categories/:categoryId", categoryController.Update)
	router.DELETE(apiV1RootPath+"/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
