package defangIPaddr

import (
	"fmt"
	"testing"
)

func defangIPaddr(address string) string {
	cutStrList, qa, head := make([]string, 0), "", 0
	for i, value := range address {
		if value == 46 {
			cutStrList = append(cutStrList, address[head:i])
			head = i + 1
		}
	}
	cutStrList = append(cutStrList, address[head:])
	for i, s := range cutStrList {
		qa += string(s)
		if i != len(cutStrList)-1 {
			qa += "[.]"
		}
	}
	return qa
}

func TestDefangIPaddr(t *testing.T) {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
