package beautifulDays

import (
	"math"
	"strconv"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	result := int32(0)
	for l := i; l <= j; l++ {
		r := Reverse(strconv.Itoa(int(l)))
		ir, _ := strconv.ParseInt(r, 10, 64)
		if int(math.Abs(float64(l-int32(ir))))%int(k) == 0 {
			result++
		}
	}
	return result
}

func Reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
