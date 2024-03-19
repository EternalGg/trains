package canBeTypedWords

import "fmt"

func canBeTypedWords(text string, brokenLetters string) int {
	strGroup := splitGraph(text)
	fmt.Println(strGroup)
	result := len(strGroup)
	brokenMap := make(map[int32]bool)
	for _, value := range brokenLetters {
		brokenMap[value] = true
	}

	for _, str := range strGroup {
		for _, value := range str {
			if brokenMap[value] {
				result--
				break
			}
		}
	}
	return result
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
	newOne := text[last:]
	result = append(result, newOne)
	return result
}
