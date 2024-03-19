package countWords

func countWords(words1 []string, words2 []string) int {
	w1Map, w1Se, w2Map, w2Se, w1Check, result := make(map[string]int), make([]string, 0), make(map[string]int), make([]string, 0), make(map[string]bool), 0
	for _, str := range words1 {
		w1Map[str]++
	}
	for _, str := range words1 {
		if w1Map[str] == 1 {
			w1Se = append(w1Se, str)
			w1Check[str] = true
		}
	}
	for _, str := range words2 {
		w2Map[str]++
	}
	for _, str := range words2 {
		if w2Map[str] == 1 {
			w2Se = append(w2Se, str)
		}
	}
	for _, s := range w2Se {
		if w1Check[s] == true {
			result++
		}
	}
	return result
}
