package kClosest

import "testing"

func TestZzz(t *testing.T) {
	list := [][]int{}
	list = append(list, []int{3, 3})
	list = append(list, []int{5, -1})
	list = append(list, []int{-2, 4})
	KKClosest(list, 2)
}
