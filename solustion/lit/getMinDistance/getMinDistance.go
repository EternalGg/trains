package getMinDistance

import "math"

func getMinDistance(nums []int, target int, start int) int {
	min := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			now := int(math.Abs(float64(i - start)))
			if min == -1 {
				min = now
			} else {
				if now < min {
					min = now
				}
			}
		}
	}
	return min
}
