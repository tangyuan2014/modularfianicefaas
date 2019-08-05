package main

import (
	"modularfianicefaas/mathoperations/Operations"
	"net/http"
)

func operation(writer http.ResponseWriter, request *http.Request) {
	operand1, ok := request.URL.Query()["operand1"]
	if !ok || len(operand1[0]) < 1 {
		panic("operand1 is missing")
	}

	operand2, ok := request.URL.Query()["operand2"]
	if !ok || len(operand1[0]) < 1 {
		panic("operand2 is missing")
	}

	operator, ok := request.URL.Query()["operator"]
	if !ok || len(operator[0]) < 1 {
		panic("operand1 is missing")
	}

	res := ""
	switch operator	[0] {
	case "addition":
		res = Operations.Addition(operand1[0], operand2[0])
	case "factorial":
		res = Operations.Factorial(operand1[0])
	case "multiplication":
		res = Operations.Multi(operand1[0], operand2[0])
	default:
		panic("No such mathmatic operation")
	}
	writer.Write([]byte("result is " + res))
}

func main() {
	http.HandleFunc("/Operations/", operation)
	http.ListenAndServe(":8080", nil)
}
