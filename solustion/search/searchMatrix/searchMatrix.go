package searchMatrix

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 1 && matrix[0][0] == target {
		return true
	}
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}

	if target >= matrix[0][0] && target <= matrix[0][len(matrix[0])-1] {
		if BinarySearch(matrix[0], target) != -1 {
			return true
		} else {
			return false
		}
	}
	if target >= matrix[len(matrix)-1][0] && target <= matrix[len(matrix)-1][len(matrix[0])-1] {
		if BinarySearch(matrix[len(matrix)-1], target) != -1 {
			return true
		} else {
			return false
		}
	}

	max, min := len(matrix)-1, 0
	for !(matrix[min][0] <= target && matrix[min][len(matrix[min])-1] >= target) {
		half := ((max - min) / 2) + min

		if matrix[half][0] > target {
			if half == max {
				break
			}
			max = half
		} else {
			if half == min {
				break
			}
			min = half
		}

	}
	if BinarySearch(matrix[min], target) == -1 {
		return false
	} else {
		return true
	}
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
