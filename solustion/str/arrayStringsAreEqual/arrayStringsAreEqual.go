package arrayStringsAreEqual

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	merge1, merge2 := "", ""
	for _, s := range word1 {
		merge1 += s
	}
	for _, s := range word2 {
		merge2 += s
	}
	if merge2 == merge1 {
		return true
	} else {
		return false
	}
}
