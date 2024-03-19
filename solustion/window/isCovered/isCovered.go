package isCovered

func isCovered(ranges [][]int, left int, right int) bool {
	uMap := make(map[int]bool)
	for i := 0; i < len(ranges); i++ {
		head, last := ranges[i][0], ranges[i][1]
		for j := head; j < last; j++ {
			uMap[j] = true
		}
	}
	for i := left; i < right; i++ {
		if !uMap[i] {
			return false
		}
	}
	return true
}
