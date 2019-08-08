package operations

import (
	"strconv"
)

func Factorial(number int64) string {
	var res int64 = 1
	var i int64

	for i = 1; i <= number; i++ {
		res *= i
	}
	return strconv.FormatInt(res, 10)
}
