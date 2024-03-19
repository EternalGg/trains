package numIdenticalPairs

import (
	"fmt"
	"testing"
)

func numIdenticalPairs(nums []int) int {
	qa := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				qa++
			}
		}
	}
	return qa
}

func TestNumIdenticalPairs(t *testing.T) {
	fmt.Println(
		numIdenticalPairs([]int{1, 1, 1, 1, 3}))
}
