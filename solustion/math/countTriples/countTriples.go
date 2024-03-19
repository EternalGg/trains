package countTriples

import "math"

func CountTriples(n int) int {
	m := make(map[int]bool)
	for i := 1; i < n; i++ {
		m[int(math.Pow(float64(i), float64(2)))] = true
	}
	result := 0
	for i := n; i > 0; i-- {
		for j := i - 1; j > 0; j-- {
			if m[int(math.Pow(float64(i), float64(2)))-int(math.Pow(float64(j), float64(2)))] {
				result++
			}
		}
	}
	return result
}
