package findThePrefixCommonArray

func findThePrefixCommonArray(A []int, B []int) []int {
	aMap, result := make(map[int]bool), make([]int, len(A))
	for i := 0; i < len(A); i++ {
		aMap[A[i]] = true
		count := 0
		for j := 0; j < i; j++ {
			if aMap[B[j]] {
				count++
			}
		}
		result[i] = count
	}
	return result
}
