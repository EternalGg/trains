package minimumSum

func MinimumSum(nums []int) int {
	reverse, front := make(map[int]int), make(map[int]int)
	min := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		front[i] = min
	}
	min = nums[len(nums)-1]

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] < min {
			min = nums[i]
		}
		reverse[i] = min
	}

	result := -1
	for i := 1; i <= len(nums)-2; i++ {
		if front[i-1] < nums[i] && reverse[i+1] < nums[i] {
			if result == -1 {
				result = nums[i] + front[i-1] + reverse[i+1]
				continue
			}
			if nums[i]+front[i-1]+reverse[i+1] < result {
				result = nums[i] + front[i-1] + reverse[i+1]
			}
		}
	}

	return result

}
