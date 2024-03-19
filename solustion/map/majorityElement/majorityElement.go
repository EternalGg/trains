package majorityElement

func majorityElement(nums []int) []int {
	m := make(map[int]int)
	mc := make(map[int]bool)
	result := []int{}

	for i := 0; i < len(nums); i++ {
		if !mc[nums[i]] {
			m[nums[i]]++
			if m[nums[i]] > len(nums)/3 {
				result = append(result, nums[i])
				mc[nums[i]] = true
			}
		}
	}
	return result
}
