package server

import (
	"errors"
	"github.com/tangyuan2014/modularfianicefaas/mathoperations/cmd/operationsd/operations"
	"net/http"
	"strconv"
)

const (
	number = "number"
)

func Operation(writer http.ResponseWriter, request *http.Request) {
	number, err := validateAndGetInput(request, number)
	if err != nil {
		logAndWriteError(writer, http.StatusBadRequest, err)
		return
	}
	res := operations.Factorial(number)
	writer.Write([]byte("result is " + res))
}

func validateAndGetInput(request *http.Request, paramKey string) (int64, error) {
	param, ok := request.URL.Query()[paramKey]
	if !ok {
		return 0, errors.New("Failed to get param value from key: " + paramKey)
	} else if len(param) != 1 {
		return 0, errors.New("Please provide one and only one integer")
	}

	if digit, err := strconv.ParseInt(param[0], 10, 64); err != nil {
		return 0, errors.New("Please provide a integer ")
	} else {
		if digit < 1 {
			return 0, errors.New("Please provide a positive integer ")
		}
		return digit, nil
	}
}

func Health(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("service is up"))
}
