package server

import (
	"github.com/tangyuan2014/modularfianicefaas/mathoperations/cmd/operationsd/operations"
	"log"
	"net/http"
)

const (
	operator = "operator"
	operand1 = "operand1"
	operand2 = "operand2"
)

func Operation(writer http.ResponseWriter, request *http.Request) {
	res := ""
	operator := requestValid(writer, request, operator)
	switch operator[0] {
	case "addition":
		operand1 := requestValid(writer, request, operand1)
		operand2 := requestValid(writer, request, operand2)
		resadd, err := operations.Addition(operand1[0], operand2[0])
		if err != nil {
			notFoundError(writer)
			return
		}
		res = resadd
	case "factorial":
		operand1 := requestValid(writer, request, operand1) //TODO operand2 condition
		resfa, err := operations.Factorial(operand1[0])
		if err != nil {
			notFoundError(writer)
			return
		}
		res = resfa
	case "multiplication":
		operand1 := requestValid(writer, request, operand1)
		operand2 := requestValid(writer, request, operand2)
		resmul, err := operations.Multi(operand1[0], operand2[0])
		if err != nil {
			notFoundError(writer)
			return
		}
		res = resmul
	default:
		log.Println() //TODO
		notFoundError(writer)
		return
	}
	writer.Write([]byte("result is " + res))
}

func requestValid(writer http.ResponseWriter, request *http.Request, r string) []string {
	num, ok := request.URL.Query()[r]
	if !ok || len(num[0]) < 1 {
		log.Println() //TODO
		notFoundError(writer)
		return nil
	}
	return num
}
