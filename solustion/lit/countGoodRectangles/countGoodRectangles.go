package countGoodRectangles

import "fmt"

func countGoodRectangles(rectangles [][]int) int {
	max, count := 0, 0

	for _, rectangle := range rectangles {
		bigger := 0
		if rectangle[0] >= rectangle[1] {
			bigger = rectangle[1]
		} else {
			bigger = rectangle[0]
		}
		fmt.Println(bigger)
		if bigger > max {
			max = bigger
			count = 1
			continue
		}
		if max == bigger {
			count++
		}
	}
	return count
}
