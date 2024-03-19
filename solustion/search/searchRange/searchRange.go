package searchRange

func searchRange(nums []int, target int) []int {
	key := BinarySearch(nums, target)
	result := make([]int, 2)
	for i := key; i >= 0; i-- {
		if nums[i] != target {
			result[0] = i + 1
		}
	}
	for i := key; i < len(nums); i++ {
		if nums[i] != target {
			result[1] = i + 1
		}
	}
	return result
}

func BinarySearch(nums []int, target int) int {
	result := Binary(nums, 0, len(nums)-1, target)
	return result
}

func Binary(nums []int, head, last, target int) int {

	if head > last {
		return -1
	}
	half := ((last - head) / 2) + head
	switch {
	case nums[head] == target:
		return head
	case nums[last] == target:
		return last
	case nums[half] == target:
		return half
		// 当前 > 目标
	case nums[half] > target:
		return Binary(nums, head, half-1, target)
		// 当前 < 目标
	case nums[half] < target:
		return Binary(nums, half+1, last, target)
	}
	return -1
}
