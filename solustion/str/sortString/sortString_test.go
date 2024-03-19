package sortString

import (
	"fmt"
	"testing"
)

func TestPrefixCount(t *testing.T) {
	fmt.Println(PrefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
}
