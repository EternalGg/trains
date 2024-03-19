package commonFactors

func commonFactors(a int, b int) int {
	qa, smaller := 1, 0
	if a > b {
		smaller = b
	} else {
		smaller = a
	}

	for i := 2; i < smaller; i++ {
		if a%i == 0 && b%i == 0 {
			qa++
		}
	}
	return qa
}
