package catAndMouse

import "math"

func catAndMouse(x int32, y int32, z int32) string {
	if math.Abs(float64(x-z)) == math.Abs(float64(y-z)) {
		return "Mouse C"
	}
	if math.Abs(float64(x-z)) > math.Abs(float64(y-z)) {
		return "Cat B"
	} else {
		return "Cat A"
	}
}
