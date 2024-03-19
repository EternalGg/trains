package monkeyMove

import (
	"fmt"
	"math"
)

func MonkeyMove(n int) int {
	now := float64(0)
	if n < 100 {
		now = math.Pow(2, float64(n)) - 2
	} else {
		now = math.Pow(2, float64(n)) - 2
	}
	fmt.Println(now)
	divide := math.Pow(10, 9) + 7
	result := int64(now) % int64(divide)
	return int(result)
}

//291172002
//	pos := make(map[int]bool)
//	result := 0
//	// pos
//	for i := 0; i < n; i++ {
//		now := (i + 1) % n
//
//		if pos[now] == true {
//			result += 2
//		} else {
//			pos[now] = true
//		}
//	}
//	for i := 0; i <= n; i++ {
//		now := (i - 1 + n) % n
//
//		if pos[now] == true {
//			result += 2
//
//		} else {
//			pos[now] = true
//		}
//	}
//	return result
//}
