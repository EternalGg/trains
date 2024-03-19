package commonChars

import "fmt"

func CommonChars(words []string) []string {
	publicMap := make(map[int32]int)

	for _, letters := range words[0] {
		publicMap[letters]++
	}

	for _, word := range words {
		globalMap := make(map[int32]int)
		for _, letter := range word {
			if publicMap[letter] > 0 && globalMap[letter] < publicMap[letter] {
				fmt.Println(globalMap, word)
				globalMap[letter]++
			}
		}
		publicMap = globalMap
	}

	result := make([]string, 0)
	for _, letter := range words[0] {
		if publicMap[letter] > 0 {
			publicMap[letter]--
			result = append(result, string(letter))
		}
	}
	return result
}
