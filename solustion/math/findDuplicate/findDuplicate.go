package findDuplicate

func FindDuplicate(nums []int) int {
	count := int64(0)
	for i := 0; i < len(nums); i++ {
		count += int64(nums[i])
	}
	cut := int(count - ArithmeticProgression(1, int64(len(nums)-1)))
	if cut > len(nums) {
		return 0
	} else {
		return 0
	}
}

func ArithmeticProgression(start, end int64) int64 {
	if end-start == 0 {
		return 1
	}
	if end-start%2 == 0 {
		return (end+start)*((end-start)/2) + (end - start)
	} else {
		return (end + start) * (((end - start) / 2) + 1)
	}
}
