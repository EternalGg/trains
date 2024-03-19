package numberOfMatches

import (
	"fmt"
	"testing"
)

func numberOfMatches(n int) int {
	qa := 0

	for true {
		if n == 1 {
			return qa
		} else {
			if n%2 == 0 {
				qa += n / 2
				n = n / 2
			} else {
				qa += n / 2
				n = (n / 2) + 1
			}
		}
	}
	return 0
}

func TestNumberOfMatches(t *testing.T) {
	fmt.Println(numberOfMatches(10))
}
