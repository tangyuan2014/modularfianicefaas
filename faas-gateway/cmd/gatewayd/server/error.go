package server

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Code    string
	Message string
}

func NotFoundError(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusNotFound)
	err := errorResponse{
		Code:    "404",
		Message: "service not found ",
	}
	response, _ := json.Marshal(err)
	writer.Write(response)
}
