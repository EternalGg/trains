package permute

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	m := make([][]int, 10, 10)
	for i := 0; i < 3; i++ {
		m = append(m, make([]int, 10))
	}
	fmt.Println(m)
}
