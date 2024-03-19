package equalSubstring

import (
	"fmt"
	"math"
)

func EqualSubstring(s string, t string, maxCost int) int {
	compareList := make([]float64, 0)

	for i := 0; i < len(s); i++ {
		now := math.Abs(float64(s[i]) - float64(t[i]))
		compareList = append(compareList, now)
	}

	max, left, right, cost := 0, 0, 0, float64(0)

	for right != len(s) {
		if cost+compareList[right] <= float64(maxCost) {
			cost += compareList[right]
			right++
		} else {
			cost -= compareList[left]
			left++
		}
		if right-left > max {
			fmt.Println(right, left)
			max = right - left
		}
	}
	fmt.Println(max)
	return max
}
