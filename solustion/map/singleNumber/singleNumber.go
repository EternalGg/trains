package singleNumber

func SingleNumber(nums []int) []int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if !m[nums[i]] {
			m[nums[i]] = true
		} else {
			delete(m, nums[i])
		}
	}
	result := make([]int, 0)
	for i, _ := range m {
		if m[i] {
			result = append(result, i)
		}
	}
	return result
}
