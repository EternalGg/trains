package checkArray

func CheckArray(nums []int, k int) bool {
	lefts, rights := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			lefts = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			rights = i
			break
		}
	}
	nums = nums[lefts : rights+1]
	if len(nums) < k {
		return false
	}
	right := k
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			right = i + k
			value := nums[i]
			for j := i; j < right; j++ {
				if nums[j] == 0 {
					return false
				} else {
					nums[j] -= value
					if nums[j] < 0 {
						return false
					}
				}
			}
		} else if i+k > len(nums) && nums[i] > 0 {
			return false
		}
	}

	return true
}
