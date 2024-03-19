package findWordsContaining

func findWordsContaining(words []string, x byte) []int {
	result := []int{}

	for i, word := range words {
		for _, value := range word {
			if value == int32(x) {
				result = append(result, i)
				break
			}
		}
	}
	return result
}
