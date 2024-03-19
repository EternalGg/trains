package squares

import "math"

func Squares(a int32, b int32) int32 {
	// Write your code here
	firstSquares := int32(math.Sqrt(float64(a)))
	result := 0
	for firstSquares*firstSquares <= b {
		if firstSquares*firstSquares >= a {
			result++
		}
		firstSquares++
	}
	return int32(result)
}
