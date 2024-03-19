package solustion

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
	newOne := text[last+1:]
	result = append(result, newOne)
	return result
}
