package maximumCount

func MaximumCount(nums []int) int {

	cut, pos := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			cut = i
			break
		}
	}
	if cut == len(nums) {
		return cut
	}
	for i := cut; i < len(nums); i++ {
		if nums[i] > 0 {
			pos = i
			break
		}
	}

	neg, poss := cut, len(nums)-pos
	if poss > neg {
		return poss
	} else {
		return neg
	}
}
