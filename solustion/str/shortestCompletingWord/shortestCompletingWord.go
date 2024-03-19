package shortestCompletingWord

func ShortestCompletingWord(licensePlate string, words []string) string {
	CompletingMap, letterList := make(map[int32]int), make([]int32, 0)
	for _, value := range licensePlate {
		if value >= 65 && value <= 90 {
			value += 32
		}
		if value >= 97 && value <= 122 {
			if CompletingMap[value] == 0 {
				letterList = append(letterList, value)
			}
			CompletingMap[value]++
		}

	}
	answerMap, min := make(map[string]int), -1
	for _, word := range words {
		tempoMap := make(map[int32]int)
		for _, letter := range word {
			if CompletingMap[letter] > 0 {
				tempoMap[letter]++
			}
		}
		for _, value := range letterList {
			if !(tempoMap[value] >= CompletingMap[value]) {
				goto here
			}
		}
		answerMap[word] = len(word)
		if min == -1 {
			min = len(word)
		} else {
			if min > len(word) {
				min = len(word)
			}
		}
	here:
	}
	for _, word := range words {
		if answerMap[word] == min {
			return word
		}
	}
	return ""
}
