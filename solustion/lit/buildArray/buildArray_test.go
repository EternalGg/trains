package buildArray

import (
	"fmt"
	"testing"
)

func buildArray(nums []int) []int {
	qa := make([]int, len(nums))

	for i, num := range nums {
		qa[i] = nums[num]
	}

	return qa
}

func TestBuildArray(t *testing.T) {
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
}
