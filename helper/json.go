package helper

import (
	"encoding/json"
	"net/http"

	"github.com/mnabil1718/go-restful-api/model/web"
)

func WriteToResponseJSON(writer http.ResponseWriter, contentType string, httpStatusCode int, response web.WebResponse) {
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
