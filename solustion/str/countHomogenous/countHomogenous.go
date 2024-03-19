package countHomogenous

func CountHomogenous(s string) int {
	count, last, lc := 0, s[0], 1
	for i := 1; i < len(s); i++ {
		if last != s[i] {
			if lc == 1 {
				count++
			} else if lc%2 == 0 {
				count += (lc + 1) * (lc / 2)
			} else {
				count += (lc + 1) * (lc / 2)
				count += (lc / 2) + 1
			}
			last = s[i]
			lc = 1
		} else {
			lc++
		}
	}

	if lc == 1 {
		count++
	} else if lc%2 == 0 {
		count += (lc + 1) * (lc / 2)
	} else {
		count += (lc + 1) * (lc / 2)
		count += (lc / 2) + 1
	}

	return count % 1000000007
}
