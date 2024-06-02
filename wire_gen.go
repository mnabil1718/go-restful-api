// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/mnabil1718/go-restful-api/app"
	"github.com/mnabil1718/go-restful-api/controller"
	"github.com/mnabil1718/go-restful-api/middleware"
	"github.com/mnabil1718/go-restful-api/repository"
	"github.com/mnabil1718/go-restful-api/service"
	"net/http"
)

// Injectors from injector.go:

func InitializeServer(dbEnv app.DBEnv, rootPath app.RootPath, addressUrl app.AddressURL, option ...validator.Option) *http.Server {
	categoryRepository := repository.NewCategoryRepository()
	db := app.NewDB(dbEnv)
	validate := validator.New(option...)
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController, rootPath)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := app.NewServer(authMiddleware, addressUrl)
	return server
}

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepository, service.NewCategoryService, controller.NewCategoryController)
