package getMoneySpent

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	result := int32(-1)
	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			if keyboards[i]+drives[j] == b {
				return b
			}
			if keyboards[i]+drives[j] < b && keyboards[i]+drives[j] > result {
				result = keyboards[i] + drives[j]
			}
		}
	}
	return result
}
