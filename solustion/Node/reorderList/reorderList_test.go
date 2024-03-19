package reorderList

import (
	"fmt"
	"testing"
)

func TestReorderList(t *testing.T) {
	list := []int{1, 2, 3}
	list = list[1:]
	fmt.Println(list)
}
