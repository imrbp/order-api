package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)
	PanicIfError(err)
}

func WriteToResponseBody(w gin.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encode := json.NewEncoder(w)
	err := encode.Encode(response)
	PanicIfError(err)
}
