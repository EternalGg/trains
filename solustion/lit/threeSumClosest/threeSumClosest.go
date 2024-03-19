package threeSumClosest

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	c, l := nums[0]+nums[1]+nums[2], len(nums)
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				count := nums[i] + nums[j] + nums[k]
				if count == target {
					return target
				}
				if count > c && target > 0 {
					break
				}

				if math.Abs(float64(count-target)) < math.Abs(float64(c-target)) {
					c = count
				}

			}
		}
	}
	return c
}
