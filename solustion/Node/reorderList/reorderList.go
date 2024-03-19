package reorderList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	for i := 0; i < len(list)/2; i++ {
		list[i].Next = list[len(list)-i-1]
	}

	for i := len(list) - 1; i > len(list)/2; i-- {
		fmt.Println(list[i].Val, list[len(list)-i].Val)
		list[i].Next = list[len(list)-i]
	}

}
