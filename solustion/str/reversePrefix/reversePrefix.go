package reversePrefix

func ReversePrefix(word string, ch byte) string {
	qa, key := "", 0

	for i, letters := range word {
		if letters == int32(ch) {
			key = i
			break
		}
	}
	for i := key; i >= 0; i-- {
		qa += string(word[i])
	}
	for i := key + 1; i < len(word); i++ {
		qa += string(word[i])
	}
	return qa
}
