package sumOfSquares

func sumOfSquares(nums []int) int {
	result, l := 0, len(nums)

	for i := 0; i < len(nums); i++ {
		if l%(i+1) == 0 {
			result += nums[i] * nums[i]
		}
	}
	return result
}
