package operations

import (
	"log"
	"strconv"
)

func Factorial(operator string) (string, error) {
	op1, err := strconv.ParseInt(operator, 10, 64)
	if err != nil {
		log.Fatal() //TODO
		return "", err
	}
	var res int64 = 1
	var i int64
	if op1 <= 0 {
		log.Fatal() //TODO
		return "", err
	} else {
		for i = 1; i <= op1; i++ {
			res *= i
		}
		return strconv.FormatInt(res, 10), nil
	}
}
