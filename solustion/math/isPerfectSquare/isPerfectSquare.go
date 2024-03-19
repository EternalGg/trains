package isPerfectSquare

import (
	"fmt"
	"math"
)

func IsPerfectSquare(num int) bool {
	s := math.Pow(float64(num), 0.5)
	str := fmt.Sprint(s)
	for _, value := range str {
		if value == 46 {
			return false
		}
	}
	return true
}
