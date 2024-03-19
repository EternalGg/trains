package arithmeticTriplets

import (
	"fmt"
	"testing"
)

func arithmeticTriplets(nums []int, diff int) int {
	qa, nmap := 0, make(map[int]bool)

	for _, num := range nums {
		nmap[num] = true
	}

	for i := 0; i < len(nums)-2; i++ {

		if nmap[nums[i]+diff] && nmap[nums[i]+(diff*2)] {
			qa++
		}
	}
	return qa
}

func TestArithmeticTriplets(t *testing.T) {
	fmt.Println(
		arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
}
