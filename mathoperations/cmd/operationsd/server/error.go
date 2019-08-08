package server

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Code    string
	Message string
}

func writeError(writer http.ResponseWriter, statusCode int, err error) {
	writer.WriteHeader(statusCode)
	response, _ := json.Marshal(errorResponse{
		Code:    "404",
		Message: err.Error(),
	})
	writer.Write(response)
}
