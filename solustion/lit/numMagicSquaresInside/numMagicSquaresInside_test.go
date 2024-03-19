package numMagicSquaresInside

import (
	"fmt"
	"testing"
)

func TestNumMagicSquaresInside(t *testing.T) {
	list := [][]int{}
	list = append(list, []int{3, 10, 2, 3, 4})
	list = append(list, []int{4, 5, 6, 8, 1})
	list = append(list, []int{8, 8, 1, 6, 8})
	list = append(list, []int{1, 3, 5, 7, 1})
	list = append(list, []int{9, 4, 9, 2, 9})
	fmt.Println(NumMagicSquaresInside(list))
}
