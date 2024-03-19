package rearrangeArray

func rearrangeArray(nums []int) []int {
	pos, nagi := make([]int, 0), make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			pos = append(pos, nums[i])
		} else {
			nagi = append(nagi, nums[i])
		}
	}
	result := make([]int, 0)
	for i := 0; i < len(pos); i++ {
		result = append(result, pos[i])
		result = append(result, nagi[i])
	}
	return result
}
