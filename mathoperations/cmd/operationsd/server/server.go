package server

import (
	"errors"
	"github.com/tangyuan2014/modularfianicefaas/mathoperations/cmd/operationsd/operations"
	"log"
	"net/http"
	"strconv"
)

const (
	number = "number"
)

func Health(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("service is up"))
}

func Operation(writer http.ResponseWriter, request *http.Request) {
	number, err := validateAndGetInput(request, number)
	if err != nil {
		writeError(writer, http.StatusBadRequest, err)
		log.Println("Input validation failed with error: " + err.Error())
		return
	}
	res := operations.Factorial(number)
	writer.Write([]byte("result is " + res))
}

func validateAndGetInput(request *http.Request, r string) (int64, error) {
	num, ok := request.URL.Query()[r]
	if !ok {
		return 0, errors.New("Failed to parse " + r)
	} else if len(num) != 1 {
		return 0, errors.New("Please provide one and only one integer")
	}

	if digit, err := strconv.ParseInt(num[0], 10, 64); err != nil {
		return 0, errors.New("Please provide a integer ")
	} else {
		if digit < 1 {
			return 0, errors.New("Please provide a positive integer ")
		}
		return digit, nil
	}
}
