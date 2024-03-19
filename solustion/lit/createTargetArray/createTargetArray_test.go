package createTargetArray

import (
	"fmt"
	"testing"
)

func TestCreateTargetArray(t *testing.T) {
	fmt.Println(
		CreateTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}
