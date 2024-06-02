package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mnabil1718/go-restful-api/controller"
	"github.com/mnabil1718/go-restful-api/exception"
)

type RootPath string // needed for dependency injection binding

func NewRouter(controller controller.CategoryController, rootPath RootPath) *httprouter.Router {
	path := string(rootPath)
	router := httprouter.New()
	router.GET(path+"/categories", controller.FindAll)
	router.POST(path+"/categories", controller.Create)
	router.GET(path+"/categories/:categoryId", controller.FindById)
	router.PUT(path+"/categories/:categoryId", controller.Update)
	router.DELETE(path+"/categories/:categoryId", controller.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
