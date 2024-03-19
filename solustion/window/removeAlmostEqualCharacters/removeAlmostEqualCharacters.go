package removeAlmostEqualCharacters

func RemoveAlmostEqualCharacters(word string) int {
	if len(word) == 2 {
		if word[0] == word[1] || word[0]-1 == word[1] || word[0]+1 == word[1] {
			return 1
		} else {
			return 0
		}
	}
	result := 0
	for i := 1; i < len(word)-1; i++ {
		if word[i] == word[i+1] || word[i]-1 == word[i+1] || word[i]+1 == word[i+1] {
			result++
			i++
			continue
		}
		if word[i] == word[i-1] || word[i]-1 == word[i-1] || word[i]+1 == word[i-1] {
			result++
			i++
			continue
		}

	}
	return result
}
