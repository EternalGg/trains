package swapPairs

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	fmt.Println("1")
	for i := 1; i < len(list); i += 2 {
		middle := list[i-1]
		list[i-1] = list[i]
		list[i] = middle
	}
	fmt.Println("2")
	for i := 0; i < len(list)-1; i++ {
		list[i].Next = list[i+1]
	}
	list[len(list)-2].Next = list[len(list)-1]
	list[len(list)-1].Next = nil
	return list[0]
}
