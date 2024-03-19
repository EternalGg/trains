package maxOperations

func maxOperations(nums []int, k int) int {
	IM, result := make(map[int]int), 0
	for i := 0; i < len(nums); i++ {
		if IM[k-nums[i]] > 0 {
			result++
			IM[k-nums[i]]--
		} else {
			IM[nums[i]]++
		}
	}
	return result
}
