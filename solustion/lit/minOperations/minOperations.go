package minOperations

func MinOperations(nums []int, k int) int {
	cut, m := k, make(map[int]bool)
	for i := 1; i <= k; i++ {
		m[i] = true
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if m[nums[i]] {
			cut--
			m[nums[i]] = false
			if cut == 0 {
				return len(nums) - i
			}
		}

	}
	return len(nums)
}
