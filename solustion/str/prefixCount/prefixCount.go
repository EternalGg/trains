package prefixCount

func PrefixCount(words []string, pref string) int {
	result := 0
	for _, word := range words {
		if len(word) < len(pref) {
			continue
		}
		for key, value := range pref {
			if word[key] != uint8(value) {
				goto here
			}
		}
		result++
	here:
	}
	return result
}
