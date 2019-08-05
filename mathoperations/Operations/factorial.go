package Operations

import "strconv"

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
