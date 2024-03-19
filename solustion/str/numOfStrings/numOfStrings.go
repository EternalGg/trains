package numOfStrings

func NumOfStrings(patterns []string, word string) int {
	wMap, count := make(map[int32][]int), 0

	for i, letter := range word {
		wMap[letter] = append(wMap[letter], i)
	}

	for _, pattern := range patterns {
		head := pattern[0]
		if wMap[int32(head)] == nil {
			continue
		}
		for _, key := range wMap[int32(head)] {
			for i, letters := range pattern {
				if key+i > len(word) {
					goto here
				}
				if letters != int32(word[key+i]) {
					goto here
				}
			}
			count++
			break
		here:
		}
	}
	return count
}
