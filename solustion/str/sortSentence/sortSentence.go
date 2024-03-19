package sortSentence

func SortSentence(s string) string {
	sList, right, result := make([]string, 10), 0, ""

	for key, letter := range s {
		if letter == 32 {
			num := s[key-1]
			bToN := num - 48
			str := s[right : key-1]
			sList[bToN] = str
			right = key + 1
		}
	}
	length := len(s)
	func(length int) {
		last := s[length-1]
		bToN := last - 48
		str := s[right : length-1]
		sList[bToN] = str
	}(length)

	result = sList[1]
	for i := 2; i < len(sList); i++ {
		if sList[i] != "" {
			result += " " + sList[i]
		}
	}
	return result
}
