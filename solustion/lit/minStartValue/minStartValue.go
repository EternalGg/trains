package minStartValue

func MinStartValue(nums []int) int {
	lastNeg := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 0 {
			lastNeg = i
			break
		}
	}
	if lastNeg == 0 && nums[lastNeg] < 0 {
		return (-nums[lastNeg]) + 1
	}
	if lastNeg == 0 && nums[lastNeg] > 0 {
		return 1
	}
	result, min := 0, 0
	for i := 0; i <= lastNeg; i++ {
		result += nums[i]
		if result < min {
			min = result
		}
	}

	return (-min) + 1

}
