package findReplaceString

import (
	"fmt"
	"testing"
)

func TestFindReplaceString(t *testing.T) {
	fmt.Println(FindReplaceString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}))
}
