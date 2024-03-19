package addMinimum

func AddMinimum(word string) int {
	result, t := 0, 0
	for i := 0; i < len(word); i++ {
		switch word[i] {
		case 'a':
			switch t {
			case 0:
				if i == 0 {
					continue
				} else {
					result += 2
				}
			case 1:
				result += 1
				t = 0
			case 2:
				t = 0
			}
		case 'b':
			switch t {
			case 0:
				if i == 0 {
					result++
				}
				t = 1
			case 1:
				result += 2
			case 2:
				result += 1
				t = 1
			}
		case 'c':
			switch t {
			case 0:
				if i == 0 {
					result += 2
				} else {
					result += 1
				}
				t = 2
			case 1:
				t = 2
			case 2:
				result += 2
			}
		}
	}
	switch t {
	case 0:
		result += 2
	case 1:
		result += 1
	}
	return result
}
