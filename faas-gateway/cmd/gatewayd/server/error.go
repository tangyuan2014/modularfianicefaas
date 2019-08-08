package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Code    string
	Message string
}

func logAndWriteError(writer http.ResponseWriter, statusCode int, err error) {
	log.Println(err.Error())
	writer.WriteHeader(statusCode)
	response, _ := json.Marshal(errorResponse{
		Code:    "404",
		Message: err.Error(),
	})
	writer.Write(response)
}
