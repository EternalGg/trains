package leastMinutes

func LeastMinutes(n int) int {
	if n < 4 {
		return n
	} else if n == 4 {
		return 3
	} else {
		pow, time := 2, 0
		for pow <= n {
			time++
			pow = 2 * pow
		}
		return time + 1
	}
}
