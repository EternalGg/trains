package processQueries

import (
	"fmt"
	"testing"
)

func TestProcessQueries(t *testing.T) {
	list := []int{3, 1, 2, 1}
	fmt.Println(ProcessQueries(list, 5))
}
