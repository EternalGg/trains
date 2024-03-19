package minOperations

import (
	"fmt"
	"math"
)

func MinOperations(n int) int {
	time := 0
	for n != 0 {
		sq := float64(0)
		for n != 0 {
			if n > 0 {
				if int(math.Pow(2, sq)) >= n {
					n -= int(math.Pow(2, sq))
					break
				}
			} else {
				if int(math.Pow(2, sq))+2 >= -n || int(math.Pow(2, sq))+1 >= -n {
					n += int(math.Pow(2, sq))
					break
				}
			}
			sq++
		}
		time++
		sq = 0
	}
	return time - 1
}

func MinOperations1(n int) int {
	time := 0
	for n != 0 {
		sq := int(math.Pow(float64(n), 0.5))
		n -= int(math.Pow(2, float64(sq)))
		fmt.Println(n)
	}
	return time
}
