package change

import (
	"fmt"
	"testing"
)

func TestChange(t *testing.T) {
	fmt.Println(Change(5, []int{1, 2, 5}))
	Change1(5, []int{1, 2, 5})
}
