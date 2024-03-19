package reverseWords

func ReverseWords(s string) string {
	result, lastBlank := "", 0
	for key, value := range s {
		if value == 32 {
			result += Reverse(s[lastBlank:key]) + " "
			lastBlank = key + 1
		}
	}
	result += Reverse(s[lastBlank:])
	return result
}

func Reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func ReverseWord(s string) string {
	words, result := NewSplitGraph(s), ""

	for i := len(words) - 1; i >= 1; i-- {
		result += words[i] + " "
	}
	result += words[0]

	return result
}

func NewSplitGraph(text string) []string {
	last, result := "", make([]string, 0)
	for i := 0; i <= len(text)-1; i++ {
		if text[i] != 32 {
			last += string(text[i])
		} else if last != "" {
			result = append(result, last)
			last = ""
		}
	}
	if last != "" {
		result = append(result, last)
	}
	return result
}
