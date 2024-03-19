package nextLargerNodes

import (
	"fmt"
	"testing"
)

func TestNextLargerNodes(t *testing.T) {
	node := new(ListNode)
	node.Val = 1
	node1 := new(ListNode)
	node1.Val = 7
	node2 := new(ListNode)
	node2.Val = 5
	node3 := new(ListNode)
	node3.Val = 1
	node4 := new(ListNode)
	node4.Val = 9
	node5 := new(ListNode)
	node6 := new(ListNode)
	node7 := new(ListNode)
	node5.Val = 2
	node6.Val = 5
	node7.Val = 1
	node.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	fmt.Println(NextLargerNodes(node))
}
