package maxDepth

import (
	"fmt"
	"testing"
)

func maxDepth(s string) int {
	now, max := 0, 0

	for _, ward := range s {
		if ward == 40 {
			now++
			if now > max {
				max = now
			}
		}
		if ward == 41 {
			now--
		}
	}
	return max
}

func TestMaxDepth(t *testing.T) {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}
