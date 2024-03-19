package addBinary

import (
	"fmt"
	"strconv"
)

func AddBinary(a string, b string) string {
	ai, _ := strconv.ParseInt(a, 2, 64)
	bi, _ := strconv.ParseInt(b, 2, 64)
	fmt.Println(ai, bi, ai+bi)
	return IntegerToBinary(ai + bi)
}

func IntegerToBinary(n int64) string {
	return strconv.FormatInt(n, 2)
}
