package lengthOfLIS

func LengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	dp, result := make([]int, 1), 0
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		max := 1

		for j := 0; j < len(dp); j++ {
			if nums[i] > nums[j] && dp[j] >= max {
				max = dp[j] + 1
			}
		}
		if nums[i] > nums[i-1] && dp[i-1] >= max {
			max = dp[i-1] + 1
		}
		dp = append(dp, max)
		if max > result {
			result = max
		}
	}
	return result
}
