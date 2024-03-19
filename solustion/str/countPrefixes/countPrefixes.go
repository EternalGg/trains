package countPrefixes

func CountPrefixes(words []string, s string) int {
	count := 0

	for _, word := range words {
		if len(word) > len(s) {
			continue
		}
		for i2, s2 := range s {
			if len(word) == i2 {
				break
			}
			if s2 != int32(word[i2]) {
				goto here
			}
		}
		count++
		continue
	here:
	}
	return count
}
