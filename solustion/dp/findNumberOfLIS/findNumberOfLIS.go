package findNumberOfLIS

func LengthOfLIS(nums []int) int {

	dp, result := make([]int, 1), 0
	dp[0] = 1
	cDp := make([]int, 1)
	for i := 1; i < len(nums); i++ {
		max, count := 1, 0

		for j := 0; j < len(dp); j++ {
			if nums[i] > nums[j] && dp[j] >= max {
				max = dp[j] + 1
				count++
			}
		}
		if nums[i] >= nums[i-1] && dp[i-1] >= max {
			max = dp[i-1] + 1
			count++
		}
		cDp = append(cDp, count)
		dp = append(dp, max)
		if max > result {
			result = max
		}
	}
	answer := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] == result-1 && cDp[i] > answer {
			answer = cDp[i]
		}
	}
	if answer == 1 {
		return len(nums)
	} else {
		return answer
	}
}
