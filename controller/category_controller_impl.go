package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mnabil1718/go-restful-api/helper"
	"github.com/mnabil1718/go-restful-api/model/web"
	"github.com/mnabil1718/go-restful-api/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{CategoryService: categoryService}
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categories := controller.CategoryService.FindAll(request.Context())
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categories,
	}
	helper.WriteToResponseJSON(writer, "application/json", http.StatusOK, response)
}
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)
	category := controller.CategoryService.FindById(request.Context(), int64(categoryId))
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   category,
	}
	helper.WriteToResponseJSON(writer, "application/json", http.StatusOK, response)
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.DecodeRequestFromJSON(request, &categoryCreateRequest)
	category := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   category,
	}
	helper.WriteToResponseJSON(writer, "application/json", http.StatusOK, response)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.DecodeRequestFromJSON(request, &categoryUpdateRequest)
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = int64(categoryId)
	category := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   category,
	}
	helper.WriteToResponseJSON(writer, "application/json", http.StatusOK, response)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)
	controller.CategoryService.Delete(request.Context(), int64(categoryId))
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}
	helper.WriteToResponseJSON(writer, "application/json", http.StatusOK, response)
}
