package double

import (
	"math"
	"strconv"
)

func IntegerToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func binaryToInteger(num int) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0 {
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}
