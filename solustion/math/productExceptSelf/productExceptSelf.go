package productExceptSelf

func productExceptSelf(nums []int) []int {
	zeroTime, zerokey, count := 0, -1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroTime++
			zerokey = i
		} else {
			count *= nums[i]
		}
	}
	result := make([]int, len(nums))
	if zeroTime == 0 {
		for i := 0; i < len(nums); i++ {
			result[i] = count / nums[i]
		}
		return result
	} else if zeroTime == 1 {
		result[zerokey] = count
		return result
	} else {
		return result
	}
}
