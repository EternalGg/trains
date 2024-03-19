package getDecimalValue

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetDecimalValue(head *ListNode) int {
	qa, str := int64(0), ""
	for head.Next != nil {
		now := strconv.Itoa(head.Val)
		str += now
	}
	fmt.Println(str)
	qa, _ = strconv.ParseInt(str, 2, 32)
	return int(qa)
}
