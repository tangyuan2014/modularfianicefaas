package operations

import (
	"log"
	"strconv"
)

func Addition(operator1 , operator2 string) (string,error) {
	op1, err := strconv.ParseInt(operator1, 10, 64)
	if err != nil {
		log.Fatal()//TODO
		return "",err
	}
	op2, err := strconv.ParseInt(operator2, 10, 64)
	if err != nil {
		log.Fatal()//TODO
		return "",err
	}
	op1 += op2
	return strconv.FormatInt(op1, 10),nil
}
