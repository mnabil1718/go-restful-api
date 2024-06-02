package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/mnabil1718/go-restful-api/app"
	"github.com/mnabil1718/go-restful-api/helper"
)

func main() {

	// ============ DEPENDENCY INJECTION
	// validator.Option is recommended by the docs: https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme
	server := InitializeServer(app.Dev, app.RootPath("/api/v1"), app.AddressURL("localhost:8080"), validator.WithRequiredStructEnabled())

	// ============ WITHOUT DEPENDENCY INJECTION
	// db := app.NewDB(app.Dev)
	// validate := validator.New()
	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)

	// apiV1Path := app.RootPath("/api/v1")
	// router := app.NewRouter(categoryController, apiV1Path)

	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: middleware.NewAuthMiddleware(router),
	// }

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
