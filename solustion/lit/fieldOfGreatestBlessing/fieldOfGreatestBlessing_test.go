package fieldOfGreatestBlessing

import (
	"fmt"
	"testing"
)

func TestFieldOfGreatestBlessing(t *testing.T) {
	list := make([][]int, 0)
	//[[],[7,5,3],[1,8,5],[5,6,3],[9,10,2],[8,4,10]]
	list = append(list, []int{4, 4, 6})
	list = append(list, []int{7, 5, 3})
	list = append(list, []int{1, 6, 2})
	list = append(list, []int{5, 6, 3})
	fmt.Println(FieldOfGreatestBlessing(list))
}
