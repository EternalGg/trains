package isGood

import "sort"

func isGood(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != i+1 {
			return false
		}
	}
	if nums[len(nums)-1] != len(nums)-1 {
		return false
	}
	return true
}
