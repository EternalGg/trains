package interpret

import (
	"fmt"
	"testing"
)

func interpret(command string) string {
	qa := ""
	for i := 0; i < len(command); i++ {
		if command[i] == 40 {
			if command[i+1] == 41 {
				qa += "o"
				i++
			} else {
				qa += "al"
				i += 3
			}
		} else {
			qa += "G"
		}
	}
	return qa
}

func TestInterpret(t *testing.T) {
	fmt.Println(interpret("G()(al)()()()"))
}
