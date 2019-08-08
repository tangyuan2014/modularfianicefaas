package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Code    int
	Message string
}

func logAndWriteError(writer http.ResponseWriter, statusCode int, err error) {
	log.Println(err.Error())
	writer.WriteHeader(statusCode)
	response, _ := json.Marshal(errorResponse{
		Code:    statusCode,
		Message: err.Error(),
	})
	writer.Write(response)
}
