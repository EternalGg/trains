package findIndices

import "math"

func FindIndices(nums []int, indexDifference int, valueDifference int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) >= float64(valueDifference) && math.Abs(float64(i-j)) >= float64(indexDifference) {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
