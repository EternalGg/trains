package chanel

import (
	"fmt"
	"testing"
)

func TestTestRotine(t *testing.T) {
	s := 10
	ap, bp := &s, &s
	fmt.Println(ap, bp)
	m := map[*int]int{}
	m[ap] = 1
	m[ap] = 2
	fmt.Println(len(m))
	ap = nil

	fmt.Println(ap, bp)
}
