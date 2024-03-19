package minTimeToVisitAllPoints

import (
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	qa := 0
	for i := 0; i < len(points)-1; i++ {
		A := math.Abs(float64(points[i][0] - points[i+1][0]))
		B := math.Abs(float64(points[i][1] - points[i+1][1]))
		if A > B {
			qa += int(A)
		} else {
			qa += int(B)
		}
	}
	return qa
}
