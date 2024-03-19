package findClosestNumber

import "math"

func FindClosestNumber(nums []int) int {
	result, abs := nums[0], math.Abs(float64(nums[0]))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		}
		if math.Abs(float64(nums[i])) <= abs {
			if math.Abs(float64(nums[i])) == abs {
				if nums[i] > 0 {
					result = nums[i]
				}
			} else {
				result = nums[i]
			}
			abs = math.Abs(float64(nums[i]))
		}
	}
	return result
}
