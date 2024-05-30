package exception

import (
	"net/http"

	"github.com/mnabil1718/go-restful-api/helper"
	"github.com/mnabil1718/go-restful-api/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	InternalServerError(writer, request, err)
}

func InternalServerError(writer http.ResponseWriter, request *http.Request, err any) {
	response := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseJSON(writer, "application/json", http.StatusInternalServerError, response)
}
