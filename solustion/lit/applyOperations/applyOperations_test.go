package applyOperations

import (
	"fmt"
	"testing"
)

func TestApplyOperations(t *testing.T) {
	list := []int{1, 2, 2, 1, 1, 0}
	fmt.Println(ApplyOperations(list))
}
