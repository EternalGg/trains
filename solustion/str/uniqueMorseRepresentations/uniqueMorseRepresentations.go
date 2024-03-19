package uniqueMorseRepresentations

func UniqueMorseRepresentations(words []string) int {
	qa, word, wMap := 0, []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}, make(map[string]bool)

	for _, woord := range words {
		now := ""
		for _, letters := range woord {
			letter := letters - 97
			now = now + word[int(letter)]
		}
		if wMap[now] == false {
			wMap[now] = true
			qa++
		}
	}
	return qa
}
