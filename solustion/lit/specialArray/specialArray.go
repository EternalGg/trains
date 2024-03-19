package specialArray

import "sort"

func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if len(nums)-i >= nums[i] {
			return len(nums) - i
		}
	}
	return -1
}
