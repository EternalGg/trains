package deleteNodee

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	result := head
	if head.Val == val {
		return head.Next
	}
	if head.Next.Val != val {
		fmt.Println(head.Val, head.Next.Val)
		head = head.Next
	}

	head.Next = head.Next.Next
	return result
}
