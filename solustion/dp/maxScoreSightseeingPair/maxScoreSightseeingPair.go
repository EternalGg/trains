package maxScoreSightseeingPair

func maxScoreSightseeingPair(values []int) int {
	max, dp := 0, 0
	dp = values[0]
	for i := 1; i < len(values); i++ {
		dp--
		if values[i] > dp {
			dp = values[i]
		}
		if dp+values[i] > max {
			max = dp + values[i]
		}
	}
	return max
}
