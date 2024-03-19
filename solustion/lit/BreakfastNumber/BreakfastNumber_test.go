package BreakfastNumber

import (
	"fmt"
	"testing"
)

func TestBreakfastNumber(t *testing.T) {
	s := []int{7, 3, 4, 3, 9, 9, 10, 8, 8, 3}
	d := []int{7, 10, 6, 7, 5, 2, 8, 4, 5, 8}
	fmt.Println(BreakfastNumber(s, d, 5))
}
