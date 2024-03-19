package countNegatives

import (
	"testing"
)

func TestFindFirstNegative(t *testing.T) {
	list := []int{1, 0, 0, -1, -1}
	FindFirstNegative(list)
}
