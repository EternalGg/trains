package canVisitAllRooms

import (
	"fmt"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	list := [][]int{}
	list = append(list, []int{1, 3})
	list = append(list, []int{3, 0, 1})
	list = append(list, []int{2})
	list = append(list, []int{0})
	fmt.Println(CanVisitAllRooms(list))
}
