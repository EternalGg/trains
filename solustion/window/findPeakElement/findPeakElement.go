package findPeakElement

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Compare(nums[0], nums[1])
	}
	if nums[0] > nums[1] && nums[0] > nums[len(nums)-1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] && nums[len(nums)-1] > nums[0] {
		return len(nums) - 1
	}
	base, left, right := 1, 0, 2
	for !(nums[base] > nums[left]) || !(nums[base] > nums[right]) {
		if right == len(nums)-1 {
			break
		}
		base++
		left++
		right++
	}
	return base
}

func Compare(a, b int) int {
	if a > b {
		return 0
	} else {
		return 1
	}
}
