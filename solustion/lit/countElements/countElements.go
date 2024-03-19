package countElements

func countElements(nums []int) int {
	max, min := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			continue
		}
		if nums[i] < min {
			min = nums[i]
			continue
		}
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		if !(nums[i] == max || nums[i] == min) {
			result++
		}
	}
	return result
}
