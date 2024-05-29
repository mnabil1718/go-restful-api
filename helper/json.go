package helper

import (
	"encoding/json"
	"net/http"

	"github.com/mnabil1718/go-restful-api/model/web"
)

func BuildWebResponse(writer http.ResponseWriter, data any, contentType string) {
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", contentType)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func DecodeRequestFromJSON(request *http.Request, requestStruct any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&requestStruct)
	PanicIfError(err)
}
