package numEquivDominoPairs

import (
	"fmt"
	"testing"
)

func TestNumEquivDominoPairs(t *testing.T) {
	matrix := [][]int{}
	matrix = append(matrix, []int{1, 2})

	matrix = append(matrix, []int{2, 1})

	matrix = append(matrix, []int{1, 2})
	fmt.Println(NumEquivDominoPairs(matrix))
}
