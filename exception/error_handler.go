package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mnabil1718/go-restful-api/helper"
	"github.com/mnabil1718/go-restful-api/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err any) bool {

	if exception, ok := err.(validator.ValidationErrors); ok {
		response := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseJSON(writer, "application/json", http.StatusBadRequest, response)
		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {

	if exception, ok := err.(NotFoundError); ok {
		response := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Message,
		}

		helper.WriteToResponseJSON(writer, "application/json", http.StatusNotFound, response)
		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {
	response := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseJSON(writer, "application/json", http.StatusInternalServerError, response)
}
