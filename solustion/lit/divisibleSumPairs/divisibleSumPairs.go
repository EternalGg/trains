package divisibleSumPairs

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	result := int32(0)
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				result++
			}
		}
	}
	return result
}
