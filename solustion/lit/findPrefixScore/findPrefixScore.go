package findPrefixScore

func findPrefixScore(nums []int) []int64 {
	max, result := nums[0], make([]int64, len(nums))
	result[0] = int64(2 * nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			result[i] = result[i-1] + int64(nums[i]*2)
			max = nums[i]
		} else {
			result[i] = result[i-1] + int64(nums[i]) + int64(max)
		}
	}
	return result
}
