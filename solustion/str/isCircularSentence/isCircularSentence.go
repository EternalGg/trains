package isCircularSentence

import "fmt"

func isCircularSentence(sentence string) bool {
	strGroup := splitGraph(sentence)
	if len(strGroup) == 1 {
		if sentence[0] == sentence[len(sentence)-1] {
			return true
		} else {
			return false
		}
	}
	fmt.Println(strGroup)
	for i, str := range strGroup {
		if i == len(strGroup)-1 {
			break
		}
		if str[len(str)-1] != strGroup[i+1][0] {
			fmt.Println("0-n-1")
			return false
		}
	}
	if strGroup[len(strGroup)-1][len(strGroup[len(strGroup)-1])-1] == strGroup[0][0] {
		return true
	} else {
		fmt.Println("last")
		return false
	}
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
