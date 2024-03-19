package minOperations

func MinOperations(logs []string) int {
	result := 0

	for _, log := range logs {
		switch len(log) {
		case 2:
			if log[0] == '.' {
				break
			} else {
				result++
			}
		case 3:
			if log[0] == '.' && log[1] == '.' {
				if result != 0 {
					result--
				}
			} else {
				result++
			}
		default:
			result++
		}
	}

	return result
}
