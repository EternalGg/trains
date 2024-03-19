package longestMountain

import (
	"fmt"
	"testing"
)

func TestLongestMountain(t *testing.T) {
	list := []int{3, 3, 1, 0, 1, 0, 1, 0, 2, 1}
	fmt.Println(LongestMountain(list))
}
