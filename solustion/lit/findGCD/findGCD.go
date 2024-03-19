package findGCD

func findGCD(nums []int) int {
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	qa := min
	for i := 0; i < min; i++ {
		if max%qa == 0 && min%qa == 0 {
			return qa
		} else {
			qa--
		}
	}
	return qa
}
