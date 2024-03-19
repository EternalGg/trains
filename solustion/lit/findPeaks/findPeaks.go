package findPeaks

func findPeaks(mountain []int) []int {
	result := []int{}
	for i := 1; i < len(mountain)-2; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			result = append(result, i)
		}
	}
	return result
}
