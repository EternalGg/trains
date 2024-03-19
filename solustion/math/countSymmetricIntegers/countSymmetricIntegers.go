package countSymmetricIntegers

import (
	"strconv"
)

func CountSymmetricIntegers(low int, high int) int {
	result := 0
	if low <= 10 {
		low = 10
	}
	if low <= 1000 && low >= 100 {
		low = 1000
	}
	for i := low; i <= high; i++ {
		str := strconv.Itoa(i)
		switch len(str) {
		case 2:
			result += 1
			i += 10
		case 3:
			i = 1000
		case 4:
			if (i/1%10 + i/10%10) == (i/100%10 + i/1000%10) {
				result++
			}
		}
	}
	return result
}
