package minCapability

import "fmt"

func MinCapability(nums []int, k int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = min(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = min(nums[i-1], nums[i-2]+nums[i])
	}
	fmt.Println("lmap")
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
