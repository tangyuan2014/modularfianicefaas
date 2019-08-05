package main

import (
	"net/http"
	"strconv"
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
		res = Addition(operand1[0], operand2[0])
	case "factorial":
		res = Factorial(operand1[0])
	case "multiplication":
		res = Multi(operand1[0], operand2[0])
	default:
		panic("No such mathmatic operation")
	}
	writer.Write([]byte("result is " + res))
}

func Addition(operator1 , operator2 string) string {
	op1, err := strconv.ParseInt(operator1, 10, 64)
	if err != nil {
		panic("no such ")
	}
	op2, err := strconv.ParseInt(operator2, 10, 64)
	if err != nil {
		panic("no such ")
	}
	op1 += op2
	return strconv.FormatInt(op1, 10)
}

func Factorial(operator string) string {
	op1, err := strconv.ParseInt(operator, 10, 64)
	if err != nil {
		panic("no such ")
	}
	var res int64 = 1
	var i int64
	if op1 <= 0 {
		panic("Factorial of negative number or zero doesn't exit")
	} else {
		for i = 1; i <= op1; i++ {
			res *= i
		}
		return strconv.FormatInt(res, 10)
	}
}

func Multi(operator1, operator2 string) string {
	op1, err := strconv.ParseInt(operator1, 10, 64)
	if err != nil {
		panic("no such ")
	}
	op2, err := strconv.ParseInt(operator2, 10, 64)
	if err != nil {
		panic("no such ")
	}
	op1 *= op2
	return strconv.FormatInt(op1, 10)
}

func main() {
	http.HandleFunc("/Operations/", operation)
	http.ListenAndServe(":8080", nil)
}
