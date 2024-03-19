package myAtoi

import (
	"math"
	"strconv"
)

func MyAtoi(s string) int {
	result := ""
	for _, value := range s {
		switch {
		case value >= 48 && value <= 57:
			result += string(value)

		case value == 32 || value == 43:
			if len(result) > 0 && value == 32 {
				goto here
			}
			continue
		case value == 45:
			if len(result) > 0 {
				return 0
			}
			result += string(value)
		default:
			goto here
		}
	}
here:
	num, _ := strconv.ParseInt(result, 10, 64)
	if float64(num) > math.Pow(float64(2), float64(31))-1 {
		return int(math.Pow(float64(2), float64(31)) - 1)
	}
	if float64(num) < -math.Pow(float64(2), float64(31))-1 {
		return int(-math.Pow(float64(2), float64(31)))
	}
	return int(num)
}
