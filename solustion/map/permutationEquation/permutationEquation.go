package permutationEquation

func permutationEquation(p []int32) []int32 {
	// Write your code here
	nMap := make(map[int]int)
	for i := 0; i < len(p); i++ {
		nMap[int(p[i])] = i
	}

	result := []int32{}
	for i := 1; i <= len(p); i++ {
		y1 := nMap[i]
		y2 := nMap[y1]

		result = append(result, int32(y2))
	}
	return result
}
