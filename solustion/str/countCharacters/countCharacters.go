package countCharacters

func CountCharacters(words []string, chars string) int {
	letterMap, result := make(map[int32]int), 0
	for _, value := range chars {
		letterMap[value]++
	}

	for _, word := range words {
		compareMap := make(map[int32]int)
		for _, value := range word {
			compareMap[value]++
			if compareMap[value] > letterMap[value] {
				goto here
			}
		}
		result += len(word)
	here:
	}
	return result
}
