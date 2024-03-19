package findLengthOfLCIS

func findLengthOfLCIS(nums []int) int {
	result, now := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			now++
			if now > result {
				result = now
			}
		} else {
			now = 1
		}
	}
	return result
}
