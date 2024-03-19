package missingNumber

import "sort"

func missingNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return -1
}
