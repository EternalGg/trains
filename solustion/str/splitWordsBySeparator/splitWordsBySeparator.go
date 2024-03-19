package splitWordsBySeparator

func splitWordsBySeparator(words []string, separator byte) []string {
	result, last := []string{}, ""

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if words[i][j] != separator {
				last += string(words[i][j])
			} else {
				if last != "" {
					result = append(result, last)
					last = ""
				}
			}
		}
		if last != "" {
			result = append(result, last)
			last = ""
		}
	}
	return result
}
