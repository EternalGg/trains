package minSubArrayLen

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	list := []int{1, 1, 1, 7}
	fmt.Println(MinSubArrayLen(7, list))
}
