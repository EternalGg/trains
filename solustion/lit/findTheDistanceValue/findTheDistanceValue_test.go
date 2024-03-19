package findTheDistanceValue

import (
	"fmt"
	"testing"
)

func TestFindTheDistanceValue(t *testing.T) {
	arr1 := []int{2, 1, 100, 3}
	arr2 := []int{-5, -2, 10, -3, 7}
	fmt.Println(FindTheDistanceValue(arr1, arr2, 6))
}
