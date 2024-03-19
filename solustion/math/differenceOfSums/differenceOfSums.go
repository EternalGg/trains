package differenceOfSums

func differenceOfSums(n int, m int) int {
	n1, n2 := 0, 0
	for i := 1; i < n; i++ {
		if n < m {
			n1 += n
		} else {
			if n%m == 0 {
				n2++
			} else {
				n1++
			}
		}
	}
	return n1 - n2
}
