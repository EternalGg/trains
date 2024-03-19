package minimumOperations

func MinimumOperations(nums []int) int {
	result, uMap := 0, make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if uMap[nums[i]] == false && nums[i] != 0 {
			uMap[nums[i]] = true
			result++
		}
	}
	return result
}
