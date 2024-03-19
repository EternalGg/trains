package supplyWagon

import (
	"fmt"
	"testing"
)

func TestSupplyWagon(t *testing.T) {
	list := []int{1, 3, 1, 5}
	fmt.Println(SupplyWagon(list))
}
