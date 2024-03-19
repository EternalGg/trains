package carPooling

func carPooling(trips [][]int, capacity int) bool {
	carMap := make(map[int]int)
	for i := 0; i < len(trips); i++ {
		for j := trips[i][1]; j < trips[i][2]-1; j++ {
			if carMap[j]+trips[i][0] > capacity {
				return false
			} else {
				carMap[j] += trips[i][0]
			}
		}
	}
	return true
}
