package alternatingSubarray

func alternatingSubarray(nums []int) int {
	max, length := 0, len(nums)
	for i := 0; i < length; i++ {
		if max >= length {
			break
		}
		now := 0
		for j := i + 1; j < length; j++ {
			if j != nums[i]+1 {
				break
			}
			now++
		}
		if now > max {
			max = now
		}
	}
	return max
}
