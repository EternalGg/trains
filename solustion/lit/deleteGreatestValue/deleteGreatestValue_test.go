package deleteGreatestValue

import (
	"fmt"
	"testing"
)

func TestDeleteGreatestValue(t *testing.T) {
	list := [][]int{}
	list = append(list, []int{1, 2, 4})

	list = append(list, []int{3, 3, 1})
	fmt.Println(DeleteGreatestValue(list))
}
