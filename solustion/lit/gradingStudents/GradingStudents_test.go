package gradingStudents

import (
	"fmt"
	"testing"
)

func TestGradingStudents(t *testing.T) {
	list := []int32{73, 67, 38, 33}
	fmt.Println(GradingStudents(list))
}
