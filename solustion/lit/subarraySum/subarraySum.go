package subarraySum

func subarraySum(nums []int, k int) int {
	result, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count+nums[i] == k || count+nums[i] > k {

		}
	}
	return result
}
