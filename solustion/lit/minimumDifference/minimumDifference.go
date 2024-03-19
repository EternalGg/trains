package minimumDifference

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	min := nums[k-1] - nums[0]
	for i := 1; i < len(nums)-k; i++ {
		fmt.Println(nums[k+i-1], nums[i])
		if nums[k+i-1]-nums[i] < min {
			min = nums[k+i-1] - nums[i]
		}
	}
	return min
}
