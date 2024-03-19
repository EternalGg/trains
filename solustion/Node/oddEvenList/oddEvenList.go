package oddEvenList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	odd, even := head, head.Next
	result, last := odd, even
	head = head.Next
	for head != nil {
		if head.Next == nil {
			break
		}
		fmt.Println(head)
		head = head.Next
		fmt.Println(head)
		odd.Next = head
		odd = odd.Next
		if head.Next == nil {
			break
		}
		head = head.Next
		even.Next = head
		even = even.Next
	}
	odd.Next = last

	return result
}
