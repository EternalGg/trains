package getNoZeroIntegers

import (
	"math"
	"strconv"
)

func GetNoZeroIntegers(n int) []int {
	result, nToStr := make([]int, 2), strconv.Itoa(n)
	for i := 0; i < len(nToStr); i++ {
		if nToStr[i] == 48 {
			result[0] += int(math.Pow(10, float64((len(nToStr)-i)-1))) * 1
			result[1] += int(math.Pow(10, float64((len(nToStr)-i)-1))) * 9
		} else {
			result[1] += int(math.Pow(10, float64((len(nToStr)-i)-1))) * int((nToStr[i] - 48))
		}
	}
	return result
}
