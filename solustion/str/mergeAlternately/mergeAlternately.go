package mergeAlternately

func MergeAlternately(word1 string, word2 string) string {
	qa := ""
	smaller, bigger := "", ""
	if len(word2) > len(word1) {
		smaller = (word1)
		bigger = (word2)
	} else {
		smaller = (word2)
		bigger = (word1)
	}

	for i := 0; i < len(smaller); i++ {
		qa += string(word1[i])
		qa += string(word2[i])
	}
	qa += bigger[len(smaller):]
	return qa
}
