package vowelStrings

func VowelStrings(words []string, queries [][]int) []int {
	lists := make([]int, len(words))
	nownum := 0
	for i, word := range words {
		if !(word[0] == 'a' || word[0] == 'e' || word[0] == 'i' || word[0] == 'o' || word[0] == 'u') {
			goto here
		}
		if word[len(word)-1] == 'a' || word[len(word)-1] == 'e' || word[len(word)-1] == 'i' || word[len(word)-1] == 'o' || word[len(word)-1] == 'u' {
			nownum++
		}
	here:
		lists[i] = nownum
	}
	result := make([]int, 0)
	for _, query := range queries {
		time := 0
		if query[1] > len(words) {
			time = len(words) - 1
		} else {
			time = query[1]
		}
		if query[0] == 0 {
			result = append(result, lists[time])
		} else {
			value := lists[time] - lists[query[0]-1]
			result = append(result, value)
		}
	}
	return result
}

func VowelStrings2(words []string, left int, right int) int {
	result := 0

	for i := left; i <= right; i++ {
		h, l := true, true
		switch words[i][0] {
		case 'a':
		case 'e':
		case 'i':
		case 'o':
		case 'u':
		default:
			h = false
		}
		if !h {
			continue
		}
		switch words[i][len(words[i])-1] {
		case 'a':
		case 'e':
		case 'i':
		case 'o':
		case 'u':
		default:
			l = false
		}
		if !l {
			continue
		}
		result++
	}

	return result
}
