package maximizeSum

func maximizeSum(nums []int, k int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	result := max
	for i := 1; i < k; i++ {
		result += result + 1
	}
	return result
}
