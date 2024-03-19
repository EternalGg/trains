package thousandSeparator

import (
	"strconv"
)

func ThousandSeparator(n int) string {

	numToStr, plus, result := strconv.Itoa(n), 0, ""
	for i := len(numToStr) - 1; i >= 0; i-- {
		plus++
		if plus == 3 && i != 0 {
			result = "." + string(numToStr[i]) + result
			plus = 0
		} else {
			result = string(numToStr[i]) + result
		}
	}

	return result
}
