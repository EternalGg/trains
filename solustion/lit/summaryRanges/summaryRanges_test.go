package summaryRanges

import (
	"fmt"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	list := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(SummaryRanges(list))
}
