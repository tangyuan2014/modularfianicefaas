package Operations

import "strconv"

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
