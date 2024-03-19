package largestLocal

import (
	"testing"
)

func largestLocal(grid [][]int) [][]int {
	length, height, qa := len(grid)-2, len(grid[0])-2, [][]int{}
	time := length * height
	for i := 0; i < length; i++ {
		qa = append(qa, []int{})
	}
	for i := 0; i < time; i++ {
		nowL, nowH := i/height, i%height
		max := 0
		for i := nowL; i < nowL+3; i++ {
			for j := nowH; j < nowH+3; j++ {
				if grid[i][j] > max {
					max = grid[i][j]
				}
			}
		}
		qa[nowL] = append(qa[nowL], max)
		max = 0
	}

	return qa
}

func TestLargestLocal(t *testing.T) {
	now := [][]int{}
	now = append(now, []int{6, 2, 2, 2, 3, 3})
	now = append(now, []int{9, 9, 8, 1, 3, 3})
	now = append(now, []int{5, 6, 2, 6, 3, 3})
	now = append(now, []int{8, 2, 6, 4, 3, 3})
	now = append(now, []int{8, 2, 6, 4, 3, 3})

	largestLocal(now)
}
