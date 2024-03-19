package maximumTripletValue

func MaximumTripletValue(nums []int) (max int64) {
	biggest, reverseM := make([]int, len(nums)), nums[len(nums)-1]
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > reverseM {
			reverseM = nums[i]
		}
		biggest[i] = reverseM
	}
	premax := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > premax {
			premax = nums[i]
		}
		if int64((premax-nums[i])*biggest[i+1]) > max {
			max = int64((premax - nums[i]) * biggest[i+1])
		}
	}
	return max
}
