package calculate

func calculate(s string) int {
	x, y := 1, 0
	for _, i2 := range s {
		if i2 == 65 {
			x = (2 * x) + y
		} else {
			y = (2 * y) + x
		}
	}
	return x + y
}
