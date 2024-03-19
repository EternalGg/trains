package addToArrayForm

import (
	"fmt"
	"testing"
)

func TestAddToArrayForm(t *testing.T) {
	list := []int{1, 2, 0, 0}
	k := 34
	fmt.Println(AddToArrayForm(list, k))
}
