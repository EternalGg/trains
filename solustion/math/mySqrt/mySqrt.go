package mySqrt

import "math"

func MySqrt(x int) int {
	return int(math.Pow(float64(x), 2))
}
