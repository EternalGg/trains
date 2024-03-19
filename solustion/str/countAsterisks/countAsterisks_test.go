package countAsterisks

import (
	"fmt"
	"testing"
)

func countAsterisks(s string) int {
	cut, needcut, qa := "", false, 0
	for _, value := range s {
		if value == 124 {
			needcut = !needcut
		}
		if needcut == false {
			cut += string(value)
		}
	}
	for _, value := range cut {
		if value == 42 {
			qa++
		}
	}
	return qa
}

func TestCountAsterisks(t *testing.T) {
	fmt.Println(countAsterisks("l|*e*et|c**o|*de|"))
}
