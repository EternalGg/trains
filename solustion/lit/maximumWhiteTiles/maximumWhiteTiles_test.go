package maximumWhiteTiles

import (
	"fmt"
	"testing"
)

func TestMaximumWhiteTiles(t *testing.T) {
	fmt.Print(MaximumWhiteTiles([][]int{{1, 5}, {10, 11}, {12, 18}, {20, 25}, {30, 32}}, 10))
}
