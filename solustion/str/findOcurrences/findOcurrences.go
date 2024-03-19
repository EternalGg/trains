package findOcurrences

import "fmt"

func FindOcurrences(text string, first string, second string) []string {
	result := make([]string, 0)
	strGroup := splitGraph(text)

	for i, str := range strGroup {
		if i == len(strGroup)-2 {
			break
		}
		fmt.Println(str)
		if str == first {

			if strGroup[i+1] == second {
				result = append(result, strGroup[i+2])
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
	newOne := text[last+1:]
	result = append(result, newOne)
	return result
}
