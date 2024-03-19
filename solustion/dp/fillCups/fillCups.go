package fillCups

func FillCups(amount []int) int {
	A, B, C, result := amount[0], amount[1], amount[2], 0
	for A != 0 || B != 0 || C != 0 {
		switch {
		case (A != 0 && B != 0) || (C != 0 && B != 0) || (A != 0 && C != 0):
			min := Compare(A, B, C)
			if min == 0 {
				B--
				C--
			} else if min == 1 {
				A--
				C--
			} else {
				B--
				A--
			}
		default:
			if A != 0 {
				result += A
			} else if B != 0 {
				result += B
			} else {
				result += C
			}
			goto here
		}
		result++
	}
here:
	return result
}

func Compare(a, b, c int) int {
	if a >= b {
		// a 比 b 大
		if b >= c {
			return 2
		} else {
			return 1
		}
	} else {
		// b 比 a 大
		if c >= a {
			return 0
		} else {
			return 2
		}
	}
}
