package checkDistances

import (
	"fmt"
	"testing"
)

func TestCheckDistances(t *testing.T) {
	fmt.Println(CheckDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
