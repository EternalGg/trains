package maximumNumberOfStringPairs

func maximumNumberOfStringPairs(words []string) int {
	m, result := make(map[string]bool), 0
	for _, word := range words {
		if !m[word] {
			m[Reverse(word)] = true
		} else {
			delete(m, word)
			result++
		}
	}

	return result
}

func Reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
