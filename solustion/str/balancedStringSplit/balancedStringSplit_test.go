package balancedStringSplit

import (
	"fmt"
	"testing"
)

func balancedStringSplit(s string) int {
	qa, rCount, lCount := 0, 0, 0
	for _, word := range s {
		if word == 76 {
			lCount++
		}
		if word == 82 {
			rCount++
		}
		if lCount == rCount {
			qa++
			lCount = 0
			rCount = 0
		}
	}

	return qa
}

func TestBalancedStringSplit(t *testing.T) {
	fmt.Println(balancedStringSplit("RLLLLRRRLR"))
}
