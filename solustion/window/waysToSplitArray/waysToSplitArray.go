package waysToSplitArray

func WaysToSplitArray(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	left, right, result := nums[0], sum-nums[0], 0
	for i := 1; i < len(nums); i++ {
		if left >= right {
			result++
		}
		left += nums[i]
		right -= nums[i]
	}
	return result
}
