package applyOperations

func ApplyOperations(nums []int) []int {

	zeroNum := 0
	result := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zeroNum++
			continue
		}
		if nums[i] != nums[i+1] {
			result = append(result, nums[i])
		} else {
			nums[i] = 2 * nums[i]
			result = append(result, nums[i])
			nums[i+1] = 0

		}
	}
	if nums[len(nums)-1] == 0 {
		zeroNum++
	}
	for i := 0; i < zeroNum; i++ {
		result = append(result, 0)
	}
	return result
}
