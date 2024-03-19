package largestPerimeter

import (
	"sort"
)

func LargestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < len(nums)-2; i++ {
		if nums[i+2]+nums[i+1] > nums[i] {
			return nums[i+2] + nums[i+1] + nums[i]
		}
	}
	return 0
}
