package divide

import "math"

func divide(a int, b int) int {
	result := 0
	if a > 0 && b > 0 {
		for a >= b {
			a -= b
			result++
		}
	}
	if a < 0 && b < 0 {
		for a <= b {
			a -= b
			result++
		}
	}
	if a < 0 && b > 0 {
		for math.Abs(float64(a)) >= float64(b) {
			a += b
			result--
		}
	}
	if a > 0 && b < 0 {
		for float64(a) >= math.Abs(float64(b)) {
			a += b
			result--
		}
	}
	if result > 2147483647 {
		return 2147483647
	}
	if result < -2147483647 {
		return -2147483647
	}
	return result
}
