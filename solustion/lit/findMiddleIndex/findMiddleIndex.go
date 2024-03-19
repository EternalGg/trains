package findMiddleIndex

func FindMiddleIndex(nums []int) int {
	left, leftCount, rightCount := 1, nums[0], 0
	for i := 2; i < len(nums); i++ {
		rightCount += nums[i]
	}
	for leftCount != rightCount {
		rightCount -= nums[left]
		leftCount += nums[left]
		left++
	}
	return left
}
