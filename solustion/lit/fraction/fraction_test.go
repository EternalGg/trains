package fraction

import (
	"fmt"
	"testing"
)

func TestFraction(t *testing.T) {
	list := []int{2147483647}
	fmt.Println(Fraction(list))
}
