package maximalSquare

import (
	"fmt"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	fmt.Println(MaximalSquare([][]byte{{'0', '1'}, {'1', '0'}}))
	fmt.Println(MaximalSquare([][]byte{{'1', 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}))
	fmt.Println("zzzz")
}
