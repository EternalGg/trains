package haveConflict

import (
	"fmt"
	"testing"
)

func TestHaveConflict(t *testing.T) {
	l1, l2 := []string{"01:15", "02:00"}, []string{"02:00", "03:00"}
	fmt.Println(HaveConflict(l1, l2))
}
