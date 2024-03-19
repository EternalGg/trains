package averageValue

import (
	"testing"
)

func averageValue(nums []int) int {
	acount, n := 0, 0
	for _, num := range nums {
		if num%3 == 0 && num%2 == 0 {
			acount += num
			n++
		}
	}
	if n == 0 {
		return 0
	}
	return acount / n
}

func TestAverageValue(t *testing.T) {
	//averageValue()
}
