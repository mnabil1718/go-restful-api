package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mnabil1718/go-restful-api/controller"
	"github.com/mnabil1718/go-restful-api/exception"
)

func NewRouter(controller controller.CategoryController, rootPath string) http.Handler {

	router := httprouter.New()
	router.GET(rootPath+"/categories", controller.FindAll)
	router.POST(rootPath+"/categories", controller.Create)
	router.GET(rootPath+"/categories/:categoryId", controller.FindById)
	router.PUT(rootPath+"/categories/:categoryId", controller.Update)
	router.DELETE(rootPath+"/categories/:categoryId", controller.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
