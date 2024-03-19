package FreqStack

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	z := Constructor()
	z.Push(5)
	z.Push(7)
	z.Push(5)
	z.Push(7)
	z.Push(4)
	z.Push(5)
	fmt.Println(z.Pop())
	fmt.Println(z.Pop())
	fmt.Println(z.Pop())
	fmt.Println(z.Pop())
	fmt.Println(z.Pop())
}
