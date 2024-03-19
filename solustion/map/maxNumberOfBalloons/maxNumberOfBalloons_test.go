package maxNumberOfBalloons

import (
	"fmt"
	"testing"
)

func maxNumberOfBalloons(text string) int {
	b, a, l, o, n, wMap, qa := int32(98), int32(97), int32(108), int32(111), int32(110), make(map[int32]int), 0
	for _, i2 := range text {
		wMap[i2]++
	}
	for true {
		if wMap[b]-1 < 0 || wMap[a]-1 < 0 || wMap[l]-2 < 0 || wMap[o]-1 < 0 || wMap[n]-1 < 0 {
			//fmt.Println(wMap[b] - 1)
			//fmt.Println(wMap[a] - 1)
			//fmt.Println(wMap[l] - 2)
			//fmt.Println(wMap[o] - 1)
			//fmt.Println(wMap[n] - 1)
			return qa
		} else {
			wMap[b] -= 1
			wMap[a] -= 1
			wMap[l] -= 2
			wMap[n] -= 1
			wMap[o] -= 1
			qa++
		}
	}
	return qa
}

func TestMaxNumberOfBalloons(t *testing.T) {
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
}
