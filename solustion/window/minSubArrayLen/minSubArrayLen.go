package minSubArrayLen

func MinSubArrayLen(target int, nums []int) int {
	min := len(nums) + 1

	for i := 0; i < len(nums); i++ {
		count := nums[i]
		if count >= target {
			return 1
		}
		for j := i + 1; j < len(nums); j++ {
			if count >= target {
				if j-i < min {
					min = j - i
					if min == 1 {
						return 1
					}
					break
				}
			} else {
				count += nums[j]
				if j-i < min && count >= target {
					min = j - i + 1
					break
				}
			}
		}
	}
	if min == len(nums)+1 {
		return 0
	} else {
		return min
	}
}
