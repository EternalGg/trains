package leftRigthDifference

import "math"

func LeftRigthDifference(nums []int) []int {
	left, right := make([]int, len(nums)), make([]int, len(nums))
	leftCount := 0
	for i, num := range nums {
		left[i] = leftCount
		leftCount += num
	}

	rightCount := 0
	for i := len(nums) - 1; i >= 0; i-- {
		right[i] = rightCount
		rightCount += nums[i]
	}

	result := make([]int, len(nums))

	for i := 0; i < len(result); i++ {
		cut := int(math.Abs(float64(left[i] - right[i])))
		result[i] = cut
	}
	return result
}
