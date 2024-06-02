//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/mnabil1718/go-restful-api/app"
	"github.com/mnabil1718/go-restful-api/controller"
	"github.com/mnabil1718/go-restful-api/middleware"
	"github.com/mnabil1718/go-restful-api/repository"
	"github.com/mnabil1718/go-restful-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

func InitializeServer(dbEnv app.DBEnv, rootPath app.RootPath, addressUrl app.AddressURL, option ...validator.Option) *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		app.NewServer,
	)
	return nil
}
