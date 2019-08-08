package errorhandling

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Code    string
	Message string
}

func NotFoundError(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	err:=errorResponse{
		Code:    "404",
		Message: "service not found ",
	}
	response,_:=json.Marshal(err)
	writer.Write(response)
	
}

func InternalError(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusInternalServerError)
	err:=errorResponse{
		Code:    "500",
		Message: "internal error",
	}
	response,_:=json.Marshal(err)
	writer.Write(response)
}
