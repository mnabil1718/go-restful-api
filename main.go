package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mnabil1718/go-restful-api/app"
	"github.com/mnabil1718/go-restful-api/controller"
	"github.com/mnabil1718/go-restful-api/helper"
	"github.com/mnabil1718/go-restful-api/middleware"
	"github.com/mnabil1718/go-restful-api/repository"
	"github.com/mnabil1718/go-restful-api/service"
)

func main() {

	db := app.NewDB(app.Dev)
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	apiV1Path := "/api/v1"
	router := app.NewRouter(categoryController, apiV1Path)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
