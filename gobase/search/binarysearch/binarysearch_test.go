package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	//odd := []int{1, 2, 3, 4, 5}
	nodd := []int{10, 13}
	fmt.Println(BinaryFindSmall(nodd, 14))

	//fmt.Println(BinaryFindSmall(nodd, 5))
}
