package searchInsert

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 4, 9}

	fmt.Println(SearchInsert(list, 6))
}
