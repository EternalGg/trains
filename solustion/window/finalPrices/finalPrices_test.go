package finalPrices

import (
	"fmt"
	"testing"
)

func TestFinalPrices(t *testing.T) {
	list := []int{8, 4, 6, 6, 2, 3}
	fmt.Println(FinalPrices(list))
}
