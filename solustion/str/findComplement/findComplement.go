package findComplement

import (
	"fmt"
	"math"
	"strconv"
)

func FindComplement(num int) int {
	str := []byte(IntegerToBinary(num))

	fmt.Println(string(str))
	for i := 0; i < len(str); i++ {
		if str[i] == 48 {
			str[i] = 49
		} else {
			str[i] = 48
		}
	}
	fmt.Println(string(str))
	bToInt, _ := strconv.ParseInt(string(str), 2, 64)

	return int(bToInt)
}

func IntegerToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func binaryToInteger(num int64) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0 {
		remainder = int(num % int64(10))
		num = num / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}
