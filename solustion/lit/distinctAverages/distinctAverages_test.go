package distinctAverages

import (
	"fmt"
	"testing"
)

func TestDistinctAverages(t *testing.T) {
	list := []int{4, 1, 4, 0, 3, 5}
	fmt.Println(DistinctAverages(list))
}
