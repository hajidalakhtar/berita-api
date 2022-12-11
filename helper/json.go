package helper

import (
	"bebasinfo/exception"
	"encoding/json"
	"net/http"
)

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	exception.PanicIfNeeded(err)

}
