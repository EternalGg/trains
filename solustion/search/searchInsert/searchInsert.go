package searchInsert

func SearchInsert(nums []int, target int) int {
	return BinaryFindSmall(nums, target) + 1
}

func BinaryFindSmall(nums []int, target int) int {
	result := BinaryFind(nums, 0, len(nums)-1, target)
	return result
}

func BinaryFind(nums []int, head, last, target int) int {
	half := ((last - head) / 2) + head
	if head > last {
		return half - 1
	}
	switch {
	case nums[head] == target:
		return head
	case nums[last] == target:
		return last
	case nums[half] == target:
		return half
		// 当前 > 目标
	case nums[half] > target:
		return BinaryFind(nums, head, half-1, target)
		// 当前 < 目标
	case nums[half] < target:
		return BinaryFind(nums, half+1, last, target)
	}
	return half
}
