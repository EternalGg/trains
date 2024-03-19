package find132pattern

func find132pattern(nums []int) bool {
	result, list := false, []int{}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] && nums[i] < nums[i+1] {
			list = append(list, nums[i])
		}
	}
	list = append(list, nums[len(nums)-1])
	length := len(list)
	for i := 1; i < length-1; i++ {
		leftSmallerLsit := []int{}
		for j := i - 1; j >= 0; j-- {
			if list[j] < list[i] {
				leftSmallerLsit = append(leftSmallerLsit, list[j])
			}
		}
		for j := i + 1; j < length; j++ {
			if list[j] < list[i] {
				for k := 0; k < len(leftSmallerLsit); k++ {
					if list[j] > leftSmallerLsit[k] {
						return true
					}
				}
			}
		}
	}
	return result
}
