package waysToMakeFair

func WaysToMakeFair(nums []int) int {
	result, odd, even := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even += nums[i]
		} else {
			odd += nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if !(i%2 == 0) {
			if even-nums[i] == odd {
				result++
			}
		} else {
			if odd-nums[i] == even {
				result++
			}
		}
	}
	return result
}
