package isPrefixOfWord

func IsPrefixOfWord(sentence string, searchWord string) int {
	strGroup := splitGraph(sentence)
	for key, str := range strGroup {
		if len(str) < len(searchWord) {
			continue
		}
		for i, letter := range searchWord {
			if int32(str[i]) != letter {
				goto here
			}
		}
		return key
	here:
	}
	return -1
}

func splitGraph(text string) []string {
	last, result := 0, make([]string, 0)
	for i := 0; i < len(text)-1; i++ {
		if text[i] == 32 {
			newOne := ""
			if last == 0 {
				newOne = text[last:i]
			} else {
				newOne = text[last+1 : i]
			}
			result = append(result, newOne)

			last = i
		}
	}
	newOne := text[last+1:]
	result = append(result, newOne)
	return result
}
