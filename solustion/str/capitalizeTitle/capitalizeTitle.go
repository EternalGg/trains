package capitalizeTitle

func CapitalizeTitle(title string) string {
	strGroup := SplitGraph(title)
	for i := 0; i < len(strGroup); i++ {
		strGroup[i] = C(strGroup[i])
	}
	result := ""
	for _, str := range strGroup {
		result += str + " "
	}

	result = result[:len(result)-1]
	return result
}

func SplitGraph(text string) []string {
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
	if len(result) == 0 {
		result = append(result, text)
	} else {
		newOne := text[last+1:]
		result = append(result, newOne)
	}

	return result
}

func C(text string) string {

	result := ""
	if len(text) <= 2 {
		for i := 0; i < len(text); i++ {
			if text[i] >= 65 && text[i] <= 90 {
				result += string(text[i] + 32)
			} else {
				result += string(text[i])
			}
		}
		return result
	}
	if text[0] >= 97 && text[0] <= 122 {
		result += string(text[0] - 32)
	} else {
		result += string(text[0])
	}
	for i := 1; i < len(text); i++ {
		if text[i] >= 65 && text[i] <= 90 {
			result += string(text[i] + 32)
		} else {
			result += string(text[i])
		}
	}
	return result
}
