package twoOutOfThree

import (
	"fmt"
	"testing"
)

func TestTwoOutOfThree(t *testing.T) {
	n1 := []int{1, 1, 3, 2}
	n2 := []int{2, 3}
	n3 := []int{3}
	result := TwoOutOfThree(n1, n2, n3)
	fmt.Println(result)
}
