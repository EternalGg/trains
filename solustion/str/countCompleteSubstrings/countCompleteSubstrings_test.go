package countCompleteSubstrings

import (
	"fmt"
	"testing"
)

func TestCountCompleteSubstrings(t *testing.T) {
	//try := map[int]int{}
	//try[9] = 1
	//fmt.Println(len(try))

	//2
	word := "gvgvvgv"
	fmt.Print(CountCompleteSubstrings(word, 2))
}
