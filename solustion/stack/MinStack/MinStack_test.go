package MinStack

import (
	"fmt"
	"testing"
)

func TestMinStack_GetMin(t *testing.T) {
	m := Constructor()
	m.Push(2)
	m.Push(0)
	m.Push(3)
	m.Push(0)
	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.GetMin())
}
