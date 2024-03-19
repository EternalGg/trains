package letterCombinations

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	matrix := make([][]string, len(digits))
	for i, i2 := range digits {
		word := []string{}
		switch i2 {
		case '2':
			word = []string{"a", "b", "c"}
		case '3':
			word = []string{"d", "e", "f"}
		case '4':
			word = []string{"g", "h", "i"}
		case '5':
			word = []string{"j", "k", "l"}
		case '6':
			word = []string{"m", "n", "o"}
		case '7':
			word = []string{"p", "q", "r", "s"}
		case '8':
			word = []string{"t", "u", "v"}
		case '9':
			word = []string{"w", "x", "y", "z"}
		}
		if i == 0 {
			matrix[i] = word
		} else {
			for j := 0; j < len(matrix[i-1]); j++ {
				for k := 0; k < len(word); k++ {
					now := matrix[i-1][j] + word[k]
					matrix[i] = append(matrix[i], now)
				}
			}
		}
	}
	return matrix[len(digits)-1]
}
