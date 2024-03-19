package totalSteps

func TotalSteps(nums []int) int {
	time, times := 0, 0
	for true {
		nums, times = KillBehind(nums)
		if times == 0 {
			return time
		}
		time++
	}
	return time
}

func KillBehind(nums []int) ([]int, int) {
	result, time, last := make([]int, 0), 0, nums[0]
	for _, num := range nums {
		// 下一个大于上一个
		if num >= last {
			result = append(result, num)
		} else {
			time++
		}
		last = num
	}
	return result, time
}
