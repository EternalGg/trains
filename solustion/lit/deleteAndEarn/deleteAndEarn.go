package deleteAndEarn

func deleteAndEarn(nums []int) int {
	list := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		list[nums[i]] += nums[i]
	}

	dp := make([]int, 10001)
	dp[0] = list[0]
	dp[1] = Max(list[0], list[1])
	for i := 2; i < len(list); i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(dp)-1]
}

func Max(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
