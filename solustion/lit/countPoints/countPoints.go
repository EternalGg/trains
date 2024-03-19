package countPoints

import "math"

func countPoints(points [][]int, queries [][]int) []int {
	result := make([]int, len(queries))

	for i, query := range queries {
		count := 0
		x2, y2, r := query[0], query[1], query[2]
		for _, point := range points {
			if math.Pow(float64(point[0]-x2), 2)+math.Pow(float64(point[1]-y2), 2) <= math.Pow(float64(r), 2) {
				count++
			}
		}
		result[i] = count
	}
	return result
}
