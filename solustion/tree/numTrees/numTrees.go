package numTrees

func numTrees(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		if i%2 == 0 {
			dp[i] = (dp[i-1] * 2) + dp[i-2]*1
		} else {
			dp[i] = (dp[i-1] * 2) + dp[i-2]*2
		}
	}
	return dp[n-1]
}
