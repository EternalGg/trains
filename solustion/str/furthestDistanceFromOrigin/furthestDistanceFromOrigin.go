package furthestDistanceFromOrigin

import "math"

func furthestDistanceFromOrigin(moves string) int {
	free, c := 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 82:
			c++
		case 76:
			//l
			c--
		default:
			free++
		}
	}
	return int(math.Abs(float64(c)) + float64(free))
}
