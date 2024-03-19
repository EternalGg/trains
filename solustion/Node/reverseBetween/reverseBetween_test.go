package reverseBetween

import "testing"

func TestReverseBetween(t *testing.T) {
	node := new(ListNode)
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node.Val = 1
	node1.Val = 2
	node2.Val = 3
	node3.Val = 4
	node4.Val = 5
	node.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	ReverseBetween(node, 1, 5)
}
