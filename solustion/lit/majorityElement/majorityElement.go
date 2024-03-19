package majorityElement

func MajorityElement(nums []int) []int {
	result, cMap, length, eMap := make([]int, 0), make(map[int]int), len(nums), make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if !eMap[nums[i]] {
			cMap[nums[i]]++
			if cMap[nums[i]] >= length/3 {
				result = append(result, nums[i])
				if len(result) >= 2 {
					copys := result
					return copys
				}
				eMap[nums[i]] = true
			}
		}
	}
	return result
}

func MajorityElement1(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i, i2 := range m {
		if i2 > len(nums) {
			return i
		}
	}
	return -1
}
