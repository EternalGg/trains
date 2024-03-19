package findWinners

import "testing"

func TestFindWinners(t *testing.T) {
	l := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	FindWinners(l)
}
