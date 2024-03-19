package matrixSum

import (
	"fmt"
	"sort"
)

func matrixSum(nums [][]int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sort.Slice(nums[i], func(j, k int) bool {
			return nums[i][j] < nums[i][k]
		})
	}
	length := len(nums)
	for i := 0; i < len(nums[0]); i++ {
		max := 0
		for j := 0; j < length; j++ {
			if nums[i][j] > max {
				max = nums[i][j]
			}
		}
		count += max
		fmt.Println(max)
	}
	return count
}
