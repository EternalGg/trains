package containsNearbyAlmostDuplicate

import (
	"fmt"
	"math"
)

func ContainsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	for i := 0; i < len(nums); i++ {
		left, right := i-indexDiff, i+indexDiff

		if left < 0 {
			left = 0
		}
		if right >= len(nums) {
			right = len(nums) - 1
		}
		fmt.Println(left, right)
		for j := left; j <= right; j++ {
			if j == i {
				continue
			}
			fmt.Println(nums[j], nums[i], j, i)
			if math.Abs(float64(nums[j]-nums[i])) < float64(valueDiff) {
				return true
			}
		}
	}
	return false
}
